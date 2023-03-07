package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/api"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/middleware"
)

func NewWsRouter() *gin.Engine {
	gin.SetMode(global.WsSetting.Mode)
	r := gin.New()

	if global.WsSetting.Mode == "release" {
		r.Use(middleware.ZapLogger(), middleware.ZapRecovery())
	} else {
		r.Use(gin.Logger(), gin.Recovery())
	}

	ws := &api.WsController{}
	r.GET("/ws", ws.Ws)

	//处理ws客户端的请求回调
	client.WsClientHandler.Register("hello", func(c *client.Client, data []byte) client.ClientResponse {
		return client.ClientResponse{
			Code: 200,
			Msg:  "to " + string(data),
			Data: gin.H{"test": "test"},
		}
	})

	go client.WsClientManage.Run()

	return r
}
