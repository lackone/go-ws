package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/utils"
	"net"
	"strconv"
)

func SendClients(from string, tos []string, msg string) error {
	for _, clientId := range tos {
		bytes, _ := client.NewOkClientRes(gin.H{
			"from":   from,
			"msg":    msg,
			"msg_id": global.SnowflakeNode.Generate().Int64(),
		}).GetByte()

		if global.IsCluster() {
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

		} else {
			//如果是单机服务，则只发送到本机
			client.WsClientManage.ClientSendMsg(bytes, clientId)
		}
	}
	return nil
}

func SendGroups(from string, groups []string, msg string) error {
	for _, group := range groups {
		bytes, _ := client.NewOkClientRes(gin.H{
			"from":   from,
			"msg":    msg,
			"msg_id": global.SnowflakeNode.Generate().Int64(),
		}).GetByte()

		if global.IsCluster() {

		} else {
			client.WsClientManage.GroupSendMsg(bytes, group)
		}
	}
	return nil
}
