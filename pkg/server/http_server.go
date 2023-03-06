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

func InitHttpServer() {
	router := routes.NewHttpRouter()

	s := http.Server{
		Addr:           fmt.Sprintf(":%d", global.HttpSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.HttpSetting.HttpReadTimeout,
		WriteTimeout:   global.HttpSetting.HttpWriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if global.HttpSetting.IsTLS {
			if err := s.ListenAndServeTLS(global.HttpSetting.TLSCertFile, global.HttpSetting.TLSKeyFile); err != nil && err != http.ErrServerClosed {
				fmt.Println("http listen error :", err.Error())
			}
		} else {
			if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Println("http listen error :", err.Error())
			}
		}
	}()

	//监控信号，实现优雅关机
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("http server shutdown ... ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("http server shutdown error :", err)
	}

	fmt.Println("http server exit ... ...")
}
