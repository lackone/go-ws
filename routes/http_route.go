package routes

import (
	"github.com/gin-gonic/gin"
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

	return r
}
