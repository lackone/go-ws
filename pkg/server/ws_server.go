package server

import (
	"context"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 初始化ws服务
func InitWsServer() {
	router := routes.NewWsRouter()

	s := http.Server{
		Addr:           fmt.Sprintf(":%d", global.WsSetting.WsPort),
		Handler:        router,
		ReadTimeout:    global.WsSetting.HttpReadTimeout,
		WriteTimeout:   global.WsSetting.HttpWriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if global.WsSetting.IsTLS {
			if err := s.ListenAndServeTLS(global.WsSetting.TLSCertFile, global.WsSetting.TLSKeyFile); err != nil && err != http.ErrServerClosed {
				fmt.Println("websocket listen error :", err.Error())
				return
			}
		} else {
			if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Println("websocket listen error :", err.Error())
				return
			}
		}
	}()

	fmt.Printf("websocket run [:%d] success \n", global.WsSetting.WsPort)

	//监控信号，实现优雅关机
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("websocket server shutdown ... ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("websocket server shutdown error :", err)
	}

	fmt.Println("websocket server exit ... ...")
}
