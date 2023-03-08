package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/pkg/app"
	"github.com/lackone/go-ws/pkg/errcode"
	"github.com/lackone/go-ws/pkg/request"
	"github.com/lackone/go-ws/pkg/service"
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

	err := service.SendClients(req.From, req.Tos, req.Msg)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

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

	err := service.SendGroups(req.From, req.Groups, req.Msg)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

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

	err := service.SendMachines(req.From, req.Ips, req.Msg)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

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

	err := service.Broadcast(req.From, req.Msg)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

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

	err := service.AddGroup(req.ClientId, req.Groups)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

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

	err := service.DelGroup(req.ClientId, req.Groups)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

	res.ToSuccess(gin.H{})
}

// 在线列表
func (h *HttpController) OnlineList(ctx *gin.Context) {
	res := app.NewResponse(ctx)

	list, err := service.OnlineList()
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
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

	list, err := service.GroupList(req.ClientId)
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

	res.ToSuccess(list)
}

// 机器列表
func (h *HttpController) MachineList(ctx *gin.Context) {
	res := app.NewResponse(ctx)

	list, err := service.MachineList()
	if err != nil {
		res.ToError(errcode.ServerError.WithDetails(err.Error()))
		return
	}

	res.ToSuccess(list)
}
