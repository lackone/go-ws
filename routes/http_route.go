package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/api"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/middleware"
)

func NewHttpRouter() *gin.Engine {
	gin.SetMode(global.HttpSetting.Mode)
	r := gin.New()

	if global.HttpSetting.Mode == "release" {
		r.Use(middleware.ZapLogger(), middleware.ZapRecovery())
	} else {
		r.Use(gin.Logger(), gin.Recovery())
	}

	r.Use(middleware.Translations())

	hc := &api.HttpController{}

	r.GET("/send_clients", hc.SendClients)
	r.GET("/send_groups", hc.SendGroups)
	r.GET("/send_machines", hc.SendMachines)
	r.GET("/broadcast", hc.Broadcast)
	r.GET("/add_group", hc.AddGroup)
	r.GET("/del_group", hc.DelGroup)
	r.GET("/online_list", hc.OnlineList)
	r.GET("/group_list", hc.GroupList)
	r.GET("/machine_list", hc.MachineList)

	return r
}
