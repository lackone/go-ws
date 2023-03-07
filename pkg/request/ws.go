package request

type SendClientsRequest struct {
	From string   `form:"from" binding:"required"`
	Tos  []string `form:"tos" binding:"required"`
	Msg  string   `form:"msg" binding:"required,min=2,max=4294967295"`
}

type SendGroupsRequest struct {
	From   string   `form:"from" binding:"required"`
	Groups []string `form:"groups" binding:"required"`
	Msg    string   `form:"msg" binding:"required,min=2,max=4294967295"`
}

type SendMachinesRequest struct {
	From string   `form:"from" binding:"required"`
	Ips  []string `form:"ips" binding:"required"`
	Msg  string   `form:"msg" binding:"required,min=2,max=4294967295"`
}

type BroadcastRequest struct {
	From string `form:"from" binding:"required"`
	Msg  string `form:"msg" binding:"required,min=2,max=4294967295"`
}

type AddGroupRequest struct {
	ClientId string   `form:"client_id" binding:"required"`
	Groups   []string `form:"groups" binding:"required"`
}

type DelGroupRequest struct {
	ClientId string   `form:"client_id" binding:"required"`
	Groups   []string `form:"groups" binding:"required"`
}

type GroupListRequest struct {
	ClientId string `form:"client_id" binding:""`
}
