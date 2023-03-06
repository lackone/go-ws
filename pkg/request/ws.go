package request

type SendClientsRequest struct {
	From int64   `form:"from" binding:"required"`
	Tos  []int64 `form:"tos" binding:"required"`
	Msg  string  `form:"msg" binding:"required,min=2,max=4294967295"`
}

type SendGroupsRequest struct {
	From   int64    `form:"from" binding:"required"`
	Groups []string `form:"groups" binding:"required"`
	Msg    string   `form:"msg" binding:"required,min=2,max=4294967295"`
}

type SendMachinesRequest struct {
	From int64    `form:"from" binding:"required"`
	Ips  []string `form:"ips" binding:"required"`
	Msg  string   `form:"msg" binding:"required,min=2,max=4294967295"`
}

type BroadcastRequest struct {
	From int64  `form:"from" binding:"required"`
	Msg  string `form:"msg" binding:"required,min=2,max=4294967295"`
}

type AddGroupRequest struct {
	ClientId int64    `form:"client_id" binding:"required"`
	Groups   []string `form:"groups" binding:"required"`
}

type DelGroupRequest struct {
	ClientId int64    `form:"client_id" binding:"required"`
	Groups   []string `form:"groups" binding:"required"`
}

type GroupListRequest struct {
	ClientId int64 `form:"client_id" binding:""`
}
