package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/utils"
	"net"
	"net/netip"
	"strconv"
)

func SendClients(from string, tos []string, msg string) error {
	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   from,
		"msg":    msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	if global.IsCluster() {
		for _, clientId := range tos {
			ip, port, _, err := utils.ParseClientId(clientId, global.WsSetting.AesKey)
			if err != nil {
				return err
			}

			if global.IsLocal(ip) {
				//如果是本地则发到本机
				client.WsClientManage.ClientSendMsg(bytes, clientId)

			} else {
				//否则，则通过grpc进行远程调用
				grpcClient, err := client.NewIMGrpcClient(net.JoinHostPort(ip, strconv.Itoa(port)))
				if err != nil {
					return err
				}
				grpcClient.SendClients(from, []string{clientId}, msg)
				grpcClient.Close()
			}
		}
	} else {
		//如果是单机服务，则只发送到本机
		client.WsClientManage.ClientSendMsg(bytes, tos...)
	}
	return nil
}

func SendGroups(from string, groups []string, msg string) error {
	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   from,
		"msg":    msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	if global.IsCluster() {
		//获取所有服务列表
		addrs := global.WsEtcdDis.ServiceList()

		if len(addrs) > 0 {
			//遍历服务，给每个服务发
			for _, addr := range addrs {
				addrPort, err := netip.ParseAddrPort(addr)
				if err != nil {
					return err
				}
				if global.IsLocal(addrPort.Addr().String()) {
					client.WsClientManage.GroupSendMsg(bytes, groups...)
				} else {
					grpcClient, err := client.NewIMGrpcClient(addr)
					if err != nil {
						return err
					}
					grpcClient.SendGroups(from, groups, msg)
					grpcClient.Close()
				}
			}
		}

	} else {
		client.WsClientManage.GroupSendMsg(bytes, groups...)
	}
	return nil
}

func SendMachines(from string, ips []string, msg string) error {
	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   from,
		"msg":    msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	if global.IsCluster() {
		//获取所有服务列表
		addrs := global.WsEtcdDis.ServiceList()

		if len(addrs) > 0 {
			for _, addr := range addrs {
				addPort, err := netip.ParseAddrPort(addr)
				if err != nil {
					return err
				}
				if global.IsLocal(addPort.Addr().String()) {
					client.WsClientManage.MachineSendMsg(bytes, ips...)
				} else {
					grpcClient, err := client.NewIMGrpcClient(addr)
					if err != nil {
						return err
					}
					grpcClient.SendMachines(from, ips, msg)
					grpcClient.Close()
				}
			}
		}
	} else {
		client.WsClientManage.MachineSendMsg(bytes, ips...)
	}
	return nil
}

func Broadcast(from string, msg string) error {
	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   from,
		"msg":    msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	if global.IsCluster() {
		//获取所有服务列表
		addrs := global.WsEtcdDis.ServiceList()

		if len(addrs) > 0 {
			for _, addr := range addrs {
				addPort, err := netip.ParseAddrPort(addr)
				if err != nil {
					return err
				}
				if global.IsLocal(addPort.Addr().String()) {
					client.WsClientManage.Broadcast(bytes)
				} else {
					grpcClient, err := client.NewIMGrpcClient(addr)
					if err != nil {
						return err
					}
					grpcClient.Broadcast(from, msg)
					grpcClient.Close()
				}
			}
		}

	} else {
		client.WsClientManage.Broadcast(bytes)
	}
	return nil
}

func AddGroup(clientId string, groups []string) error {
	if global.IsCluster() {
		ip, port, _, err := utils.ParseClientId(clientId, global.WsSetting.AesKey)
		if err != nil {
			return err
		}

		if global.IsLocal(ip) {
			getClient, ok := client.WsClientManage.GetClient(clientId)
			if !ok {
				return errors.New("客户端未找到")
			}
			client.WsClientManage.AddGroupByClient(getClient, groups...)

		} else {
			grpcClient, err := client.NewIMGrpcClient(net.JoinHostPort(ip, strconv.Itoa(port)))
			if err != nil {
				return err
			}
			grpcClient.AddGroup(clientId, groups)
			grpcClient.Close()
		}
	} else {
		getClient, ok := client.WsClientManage.GetClient(clientId)
		if !ok {
			return errors.New("客户端未找到")
		}
		client.WsClientManage.AddGroupByClient(getClient, groups...)
	}
	return nil
}

func DelGroup(clientId string, groups []string) error {
	if global.IsCluster() {
		ip, port, _, err := utils.ParseClientId(clientId, global.WsSetting.AesKey)
		if err != nil {
			return err
		}

		if global.IsLocal(ip) {
			getClient, ok := client.WsClientManage.GetClient(clientId)
			if !ok {
				return errors.New("客户端未找到")
			}
			client.WsClientManage.DelGroupByClient(getClient, groups...)

		} else {
			grpcClient, err := client.NewIMGrpcClient(net.JoinHostPort(ip, strconv.Itoa(port)))
			if err != nil {
				return err
			}
			grpcClient.DelGroup(clientId, groups)
			grpcClient.Close()
		}
	} else {
		getClient, ok := client.WsClientManage.GetClient(clientId)
		if !ok {
			return errors.New("客户端未找到")
		}
		client.WsClientManage.DelGroupByClient(getClient, groups...)
	}
	return nil
}

func OnlineList() (map[string]any, error) {
	list := gin.H{}

	if global.IsCluster() {
		//获取所有服务列表
		addrs := global.WsEtcdDis.ServiceList()

		if len(addrs) > 0 {
			for _, addr := range addrs {
				addPort, err := netip.ParseAddrPort(addr)
				if err != nil {
					return nil, err
				}
				if global.IsLocal(addPort.Addr().String()) {
					allClient := client.WsClientManage.AllClient()
					if len(allClient) > 0 {
						for _, c := range allClient {
							id := c.GetID()
							list[id] = gin.H{
								"ip": c.GetIP(),
								"id": id,
							}
						}
					}
				} else {
					grpcClient, err := client.NewIMGrpcClient(addr)
					if err != nil {
						return nil, err
					}
					res, err := grpcClient.OnlineList()
					if err != nil {
						return nil, err
					}
					allClient := gin.H{}
					err = json.Unmarshal(res.Data, &allClient)
					if err != nil {
						return nil, err
					}
					if len(allClient) > 0 {
						for k, c := range allClient {
							list[k] = c
						}
					}
					grpcClient.Close()
				}
			}
		}
	} else {
		allClient := client.WsClientManage.AllClient()
		if len(allClient) > 0 {
			for _, c := range allClient {
				id := c.GetID()
				list[id] = gin.H{
					"ip": c.GetIP(),
					"id": id,
				}
			}
		}
	}

	return list, nil
}

func GroupList(clientId string) ([]string, error) {
	list := make([]string, 0)

	if global.IsCluster() {
		if len(clientId) > 0 {
			ip, port, _, err := utils.ParseClientId(clientId, global.WsSetting.AesKey)
			if err != nil {
				return nil, err
			}

			if global.IsLocal(ip) {
				getClient, ok := client.WsClientManage.GetClient(clientId)
				if !ok {
					return nil, errors.New("客户端未找到")
				}
				list = getClient.GroupList()
			} else {
				grpcClient, err := client.NewIMGrpcClient(net.JoinHostPort(ip, strconv.Itoa(port)))
				if err != nil {
					return nil, err
				}
				res, err := grpcClient.GroupList(clientId)
				if err != nil {
					return nil, err
				}
				err = json.Unmarshal(res.Data, &list)
				if err != nil {
					return nil, err
				}
				grpcClient.Close()
			}
		} else {
			//获取所有服务列表
			addrs := global.WsEtcdDis.ServiceList()

			if len(addrs) > 0 {
				for _, addr := range addrs {
					addPort, err := netip.ParseAddrPort(addr)
					if err != nil {
						return nil, err
					}
					if global.IsLocal(addPort.Addr().String()) {
						list = append(list, client.WsClientManage.GroupList()...)
					} else {
						grpcClient, err := client.NewIMGrpcClient(addr)
						if err != nil {
							return nil, err
						}
						res, err := grpcClient.GroupList("")
						if err != nil {
							return nil, err
						}
						groupList := make([]string, 0)
						err = json.Unmarshal(res.Data, &groupList)
						if err != nil {
							return nil, err
						}
						list = append(list, groupList...)
						grpcClient.Close()
					}
				}
			}
		}
	} else {
		if len(clientId) > 0 {
			getClient, ok := client.WsClientManage.GetClient(clientId)
			if !ok {
				return nil, errors.New("客户端未找到")
			}
			list = getClient.GroupList()
		} else {
			list = client.WsClientManage.GroupList()
		}
	}

	return list, nil
}

func MachineList() ([]string, error) {
	list := make([]string, 0)

	if global.IsCluster() {
		//获取所有服务列表
		addrs := global.WsEtcdDis.ServiceList()

		if len(addrs) > 0 {
			for _, addr := range addrs {
				addPort, err := netip.ParseAddrPort(addr)
				if err != nil {
					return nil, err
				}
				if global.IsLocal(addPort.Addr().String()) {
					list = append(list, client.WsClientManage.MachineList()...)
				} else {
					grpcClient, err := client.NewIMGrpcClient(addr)
					if err != nil {
						return nil, err
					}
					res, err := grpcClient.MachineList()
					if err != nil {
						return nil, err
					}
					machineList := make([]string, 0)
					err = json.Unmarshal(res.Data, &machineList)
					if err != nil {
						return nil, err
					}
					list = append(list, machineList...)
					grpcClient.Close()
				}
			}
		}
	} else {
		list = client.WsClientManage.MachineList()
	}

	return list, nil
}
