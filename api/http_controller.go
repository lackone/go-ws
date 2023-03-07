package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/pkg/app"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/errcode"
	"github.com/lackone/go-ws/pkg/request"
	"github.com/lackone/go-ws/pkg/service"
	"net"
	"strconv"
)

type HttpController struct {
}

// 给多个客户端发消息
func (h *HttpController) SendClients(ctx *gin.Context) {
	req := request.SendClientsRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	service.SendClients(req.From, req.Tos, req.Msg)

	res.ToSuccess(gin.H{})
}

// 给组发消息
func (h *HttpController) SendGroups(ctx *gin.Context) {
	req := request.SendGroupsRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	bytes, _ := client.NewClientResponse(200, req.Msg, gin.H{"from": req.From, "msg": req.Msg}).GetByte()

	client.WsClientManage.GroupSendMsg(bytes, req.Groups...)

	res.ToSuccess(gin.H{})
}

// 给系统发消息
func (h *HttpController) SendMachines(ctx *gin.Context) {
	req := request.SendMachinesRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	bytes, _ := client.NewClientResponse(200, req.Msg, gin.H{"from": req.From, "msg": req.Msg}).GetByte()

	client.WsClientManage.MachineSendMsg(bytes, req.Ips...)

	res.ToSuccess(gin.H{})
}

// 全局广播
func (h *HttpController) Broadcast(ctx *gin.Context) {
	req := request.BroadcastRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	bytes, _ := client.NewClientResponse(200, req.Msg, gin.H{"from": req.From, "msg": req.Msg}).GetByte()

	client.WsClientManage.Broadcast(bytes)

	res.ToSuccess(gin.H{})
}

// 加入组
func (h *HttpController) AddGroup(ctx *gin.Context) {
	req := request.AddGroupRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	getClient, ok := client.WsClientManage.GetClient(req.ClientId)
	if !ok {
		res.ToError(errcode.WsClientIdNotFound)
		return
	}

	client.WsClientManage.AddGroupByClient(getClient, req.Groups...)

	res.ToSuccess(gin.H{})
}

// 退出组
func (h *HttpController) DelGroup(ctx *gin.Context) {
	req := request.DelGroupRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	getClient, ok := client.WsClientManage.GetClient(req.ClientId)
	if !ok {
		res.ToError(errcode.WsClientIdNotFound)
		return
	}

	client.WsClientManage.DelGroupByClient(getClient, req.Groups...)

	res.ToSuccess(gin.H{})
}

// 在线列表
func (h *HttpController) OnlineList(ctx *gin.Context) {
	res := app.NewResponse(ctx)

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

	res.ToSuccess(list)
}

// 组列表
func (h *HttpController) GroupList(ctx *gin.Context) {
	req := request.GroupListRequest{}
	res := app.NewResponse(ctx)
	valid, errors := app.BindAndValid(ctx, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToError(errs)
		return
	}

	var list []string

	if len(req.ClientId) > 0 {
		getClient, ok := client.WsClientManage.GetClient(req.ClientId)
		if !ok {
			res.ToError(errcode.WsClientIdNotFound)
			return
		}
		list = getClient.GroupList()
	} else {
		list = client.WsClientManage.GroupList()
	}

	res.ToSuccess(list)
}

// 机器列表
func (h *HttpController) MachineList(ctx *gin.Context) {
	res := app.NewResponse(ctx)

	machines := client.WsClientManage.GetMachines()

	res.ToSuccess(machines)
}
