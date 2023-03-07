package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/proto/im"
	"net"
	"strconv"
)

type IMService struct {
	im.UnimplementedIMServiceServer
}

func (i *IMService) SendClients(ctx context.Context, req *im.SendClientsReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.From == "" || len(req.Tos) == 0 || req.Msg == "" {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   req.From,
		"msg":    req.Msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	client.WsClientManage.ClientSendMsg(bytes, req.Tos...)

	return res, nil
}

func (i *IMService) SendGroups(ctx context.Context, req *im.SendGroupsReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.From == "" || len(req.Groups) == 0 || req.Msg == "" {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   req.From,
		"msg":    req.Msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	client.WsClientManage.GroupSendMsg(bytes, req.Groups...)

	return res, nil
}

func (i *IMService) SendMachines(ctx context.Context, req *im.SendMachinesReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.From == "" || len(req.Ips) == 0 || req.Msg == "" {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   req.From,
		"msg":    req.Msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	client.WsClientManage.MachineSendMsg(bytes, req.Ips...)

	return res, nil
}

func (i *IMService) Broadcast(ctx context.Context, req *im.BroadcastReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.From == "" || req.Msg == "" {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	bytes, _ := client.NewOkClientRes(gin.H{
		"from":   req.From,
		"msg":    req.Msg,
		"msg_id": global.SnowflakeNode.Generate().Int64(),
	}).GetByte()

	client.WsClientManage.Broadcast(bytes)

	return res, nil
}

func (i *IMService) AddGroup(ctx context.Context, req *im.AddGroupReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.ClientId == "" || len(req.Groups) == 0 {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	getClient, ok := client.WsClientManage.GetClient(req.ClientId)
	if !ok {
		res.Code = 500
		res.Msg = "客户端未找到"
		return res, nil
	}

	client.WsClientManage.AddGroupByClient(getClient, req.Groups...)

	return res, nil
}

func (i *IMService) DelGroup(ctx context.Context, req *im.DelGroupReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	if req.ClientId == "" || len(req.Groups) == 0 {
		res.Code = 500
		res.Msg = "参数不完整"
		return res, nil
	}

	getClient, ok := client.WsClientManage.GetClient(req.ClientId)
	if !ok {
		res.Code = 500
		res.Msg = "客户端未找到"
		return res, nil
	}

	client.WsClientManage.DelGroupByClient(getClient, req.Groups...)

	return res, nil
}

func (i *IMService) OnlineList(ctx context.Context, req *im.OnlineListReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	allClient := client.WsClientManage.AllClient()

	list := gin.H{}

	if len(allClient) > 0 {
		for _, c := range allClient {
			id := c.GetID()
			list[id] = gin.H{
				"addr": net.JoinHostPort(c.GetIP(), strconv.Itoa(c.GetPort())),
				"id":   id,
			}
		}
	}

	res.Data, _ = json.Marshal(list)

	return res, nil
}

func (i *IMService) GroupList(ctx context.Context, req *im.GroupListReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	var list []string

	if len(req.ClientId) > 0 {
		getClient, ok := client.WsClientManage.GetClient(req.ClientId)
		if !ok {
			res.Code = 500
			res.Msg = "客户端未找到"
			return res, nil
		}
		list = getClient.GroupList()
	} else {
		list = client.WsClientManage.GroupList()
	}

	res.Data, _ = json.Marshal(list)

	return res, nil
}

func (i *IMService) MachineList(ctx context.Context, req *im.MachineListReq) (*im.CommonRes, error) {
	res := &im.CommonRes{
		Code: 200,
		Msg:  "成功",
	}

	machines := client.WsClientManage.GetMachines()

	res.Data, _ = json.Marshal(machines)

	return res, nil
}
