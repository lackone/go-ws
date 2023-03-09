package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/app"
	"github.com/lackone/go-ws/pkg/client"
	"github.com/lackone/go-ws/pkg/errcode"
	"github.com/lackone/go-ws/pkg/utils"
	"time"
)

type WsController struct {
}

func (w *WsController) Ws(ctx *gin.Context) {
	res := app.NewResponse(ctx)

	conn, err := global.WsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		res.ToError(errcode.WsUpgradeError.WithDetails(err.Error()))
		return
	}

	ip := utils.RemoteIp(ctx.Request)

	fmt.Printf("client[%s] connect success ...\n", ip)

	//生成雪花ID
	snowflakeId := global.SnowflakeNode.Generate().Int64()
	//生成客户端ID
	clientId := utils.GenerateClientId(global.LocalIP, global.GrpcSetting.GrpcPort, snowflakeId, global.WsSetting.AesKey)

	//创建客户端
	wsClient := client.NewClient(clientId, conn, client.WsClientManage)
	//向客户端发送连接消息
	wsClient.SendCommonMsg(200, "connect success", gin.H{"client_id": clientId, "connect_time": time.Now().Format(time.RFC3339)})

	//添加客户端
	client.WsClientManage.JoinClient(wsClient)

	go wsClient.ReadLoop()
	go wsClient.WriteLoop()
}
