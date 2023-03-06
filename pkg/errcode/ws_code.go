package errcode

var (
	WsUpgradeError = NewError(20000000, "ws协议升级错误")
)
