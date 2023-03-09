package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"sync"
)

var WsClientManage = NewClientManage()

type ClientManage struct {
	clients        map[string]*Client             //所有的客户端
	clientsLock    sync.RWMutex                   //客户端读写锁
	broadcast      chan []byte                    //广播通道
	connectChan    chan *Client                   //连接通道
	disconnectChan chan *Client                   //断开通道
	groups         map[string]map[string]struct{} //同一组下面有哪些客户端
	groupsLock     sync.RWMutex                   //组锁
	machines       map[string]map[string]struct{} //同一IP下有哪些客户端
	machinesLock   sync.RWMutex                   //系统锁
}

func NewClientManage() *ClientManage {
	return &ClientManage{
		clients:        make(map[string]*Client),
		clientsLock:    sync.RWMutex{},
		broadcast:      make(chan []byte, 256),
		connectChan:    make(chan *Client),
		disconnectChan: make(chan *Client),
		groups:         make(map[string]map[string]struct{}),
		groupsLock:     sync.RWMutex{},
		machines:       make(map[string]map[string]struct{}),
		machinesLock:   sync.RWMutex{},
	}
}

func (m *ClientManage) Run() {
	for {
		select {
		case c, ok := <-m.connectChan:
			if !ok {
				return
			}
			m.AddClient(c)
			m.AddMachineByClient(c)
		case c, ok := <-m.disconnectChan:
			if !ok {
				return
			}
			m.DelClient(c)
			m.DelAllGroupByClient(c)
			m.DelMachineByClient(c)
		case msg, ok := <-m.broadcast:
			if !ok {
				return
			}
			m.Broadcast(msg)
		}
	}
}

// 客户端添加组
func (m *ClientManage) AddGroupByClient(c *Client, groups ...string) {
	m.groupsLock.Lock()
	defer m.groupsLock.Unlock()

	if len(groups) > 0 {
		for _, group := range groups {
			if _, ok := m.groups[group]; !ok {
				m.groups[group] = make(map[string]struct{})
			}
			m.groups[group][c.id] = struct{}{}
			c.AddGroup(group)
		}
	}
}

// 把客户端从所有组中删除
func (m *ClientManage) DelAllGroupByClient(c *Client) {
	m.groupsLock.Lock()
	defer m.groupsLock.Unlock()

	groupList := c.GroupList()

	if len(groupList) > 0 {
		for _, group := range groupList {
			if _, ok := m.groups[group]; ok {
				delete(m.groups[group], c.id)
				c.DelGroup(group)
			}
		}
	}
}

// 把客户端从组中删除
func (m *ClientManage) DelGroupByClient(c *Client, groups ...string) {
	m.groupsLock.Lock()
	defer m.groupsLock.Unlock()

	if len(groups) > 0 {
		for _, group := range groups {
			if _, ok := m.groups[group]; ok {
				delete(m.groups[group], c.id)
				c.DelGroup(group)
			}
		}
	}
}

// 客户端从系统下添加
func (m *ClientManage) AddMachineByClient(c *Client) {
	m.machinesLock.Lock()
	defer m.machinesLock.Unlock()

	ip := c.GetIP()

	if _, ok := m.machines[ip]; !ok {
		m.machines[ip] = make(map[string]struct{})
	}

	m.machines[ip][c.id] = struct{}{}

	val, _ := json.Marshal(gin.H{
		"id":          c.GetID(),
		"ip":          c.GetIP(),
		"connectTime": c.GetConnectTime(),
	})
	global.EtcdKV.Put(global.ETCD_WS_MACHINES+"/"+global.LocalIP+"/"+ip+"/"+c.GetID(), string(val))
}

// 把客户端从系统下删除
func (m *ClientManage) DelMachineByClient(c *Client) {
	m.machinesLock.Lock()
	defer m.machinesLock.Unlock()

	ip := c.GetIP()

	if len(ip) > 0 {
		if _, ok := m.machines[ip]; ok {
			delete(m.machines[ip], c.id)
		}

		global.EtcdKV.Del(global.ETCD_WS_MACHINES + "/" + global.LocalIP + "/" + ip + "/" + c.GetID())
	}
}

// 广播
func (m *ClientManage) Broadcast(msg []byte) {
	if len(m.clients) == 0 {
		return
	}

	for _, c := range m.clients {
		c.SendMsg(msg)
	}
}

// 给组发送消息
func (m *ClientManage) GroupSendMsg(msg []byte, groups ...string) {
	if len(groups) > 0 {
		for _, group := range groups {
			if _, ok := m.groups[group]; ok {
				for cid := range m.groups[group] {
					if _, yes := m.clients[cid]; yes {
						m.clients[cid].SendMsg(msg)
					}
				}
			}
		}
	}
}

// 给机器发送消息
func (m *ClientManage) MachineSendMsg(msg []byte, ips ...string) {
	if len(ips) > 0 {
		for _, ip := range ips {
			if _, ok := m.machines[ip]; ok {
				for cid := range m.machines[ip] {
					if _, yes := m.clients[cid]; yes {
						m.clients[cid].SendMsg(msg)
					}
				}
			}
		}
	}
}

// 给多个客户端发消息
func (m *ClientManage) ClientSendMsg(msg []byte, clientIds ...string) {
	if len(clientIds) > 0 {
		for _, id := range clientIds {
			if _, ok := m.clients[id]; ok {
				m.clients[id].SendMsg(msg)
			}
		}
	}
}

// 加入
func (m *ClientManage) JoinClient(c *Client) {
	m.connectChan <- c
}

// 离开
func (m *ClientManage) LeaveClient(c *Client) {
	m.disconnectChan <- c
}

// 添加客户端
func (m *ClientManage) AddClient(c *Client) {
	m.clientsLock.Lock()
	defer m.clientsLock.Unlock()
	m.clients[c.id] = c

	val, _ := json.Marshal(gin.H{
		"id":          c.GetID(),
		"ip":          c.GetIP(),
		"connectTime": c.GetConnectTime(),
	})
	global.EtcdKV.Put(global.ETCD_WS_ACCOUNTS+"/"+global.LocalIP+"/"+c.id, string(val))
}

// 所有客户端
func (m *ClientManage) AllClient() map[string]*Client {
	m.clientsLock.RLock()
	defer m.clientsLock.RUnlock()
	return m.clients
}

// 客户端数量
func (m *ClientManage) ClientCount() int {
	m.clientsLock.RLock()
	defer m.clientsLock.RUnlock()
	return len(m.clients)
}

// 删除客户端
func (m *ClientManage) DelClient(c *Client) {
	m.clientsLock.Lock()
	defer m.clientsLock.Unlock()
	delete(m.clients, c.id)
	close(c.send)

	global.EtcdKV.Del(global.ETCD_WS_ACCOUNTS + "/" + global.LocalIP + "/" + c.id)
}

// 获取客户端
func (m *ClientManage) GetClient(id string) (*Client, bool) {
	m.clientsLock.RLock()
	defer m.clientsLock.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}

// 组列表
func (m *ClientManage) GroupList() []string {
	m.groupsLock.RLock()
	defer m.groupsLock.RUnlock()
	list := make([]string, 0)
	if len(m.groups) > 0 {
		for k := range m.groups {
			list = append(list, k)
		}
	}
	return list
}

// 获取机器列表
func (m *ClientManage) MachineList() []string {
	m.machinesLock.RLock()
	defer m.machinesLock.RUnlock()
	list := make([]string, 0)
	if len(m.machines) > 0 {
		for k := range m.machines {
			list = append(list, k)
		}
	}
	return list
}
