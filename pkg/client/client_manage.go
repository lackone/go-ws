package client

import (
	"sync"
)

var WsClientManage = NewClientManage()

type ClientManage struct {
	clients        map[int64]*Client  //所有的客户端
	clientsLock    sync.RWMutex       //客户端读写锁
	broadcast      chan []byte        //广播通道
	connectChan    chan *Client       //连接通道
	disconnectChan chan *Client       //断开通道
	groups         map[string][]int64 //同一组下面有哪些客户端
	groupsLock     sync.RWMutex       //组锁
	systems        map[string][]int64 //同一系统下有哪些客户端
	systemsLock    sync.RWMutex       //系统锁
}

func NewClientManage() *ClientManage {
	return &ClientManage{
		clients:        make(map[int64]*Client),
		clientsLock:    sync.RWMutex{},
		broadcast:      make(chan []byte),
		connectChan:    make(chan *Client),
		disconnectChan: make(chan *Client),
		groups:         make(map[string][]int64),
		groupsLock:     sync.RWMutex{},
		systems:        make(map[string][]int64),
		systemsLock:    sync.RWMutex{},
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
		case c, ok := <-m.disconnectChan:
			if !ok {
				return
			}
			m.DelClient(c)
		case msg, ok := <-m.broadcast:
			if !ok {
				return
			}
			m.Broadcast(msg)
		}
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

// 添加客户端
func (m *ClientManage) AddClient(c *Client) {
	m.clientsLock.Lock()
	defer m.clientsLock.Unlock()
	m.clients[c.id] = c
}

// 所有客户端
func (m *ClientManage) AllClient() map[int64]*Client {
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
}
