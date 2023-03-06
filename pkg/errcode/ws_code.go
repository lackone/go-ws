package errcode

var (
	WsUpgradeError     = NewError(20000000, "ws协议升级错误")
	WsClientIdNotFound = NewError(20000001, "客户端ID未找到")
)
