package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/logger"
	"github.com/lackone/go-ws/pkg/setting"
	"github.com/lackone/go-ws/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	conf string //配置文件路径
)

func InitFlag() {
	flag.StringVar(&conf, "conf", "configs/", "配置文件路径")
	flag.Parse()
}

func init() {
	//初始化命令行参数
	InitFlag()

	//初始化配置
	setting.InitSetting(conf)

	//初始化日志
	logger.InitLogger()
}

func main() {
	router := routes.NewRouter()
	s := http.Server{
		Addr:           fmt.Sprintf(":%d", global.WsSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.WsSetting.HttpReadTimeout,
		WriteTimeout:   global.WsSetting.HttpWriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server listen err", err)
		}
	}()

	//监控SIGINT和SIGTERM信号，实现优雅关机
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln("shutdown server err", err)
	}

	log.Println("server exit...")
}
