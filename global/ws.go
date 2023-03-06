package global

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	WsUpgrader *websocket.Upgrader
)

// 获取ws协议升级默认配置
func InitWsUpgrader() error {
	if WsUpgrader == nil {
		WsUpgrader = &websocket.Upgrader{
			ReadBufferSize:  WsSetting.ReadBufferSize,
			WriteBufferSize: WsSetting.WriteBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	}
	return nil
}
