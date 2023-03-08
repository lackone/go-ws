package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/server"
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
	global.InitSetting(conf)
	//初始化日志
	global.InitLogger()
	//初始化ws升级协议
	global.InitWsUpgrader()
	//初始化雪化算法
	global.InitSnowflakeNode()
	//初始化主机IP
	global.InitLocalIP()
	//初始化etcd客户端
	global.InitEtcdClient()
}

func main() {
	//启动grpc
	grpc := server.NewGRPCServer()
	go grpc.Run()

	//启动ws
	ws := server.NewWSServer()
	go ws.Run()

	//启动http
	http := server.NewHttpServer()
	go http.Run()

	//监控信号，实现优雅关机
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("shutdown ... ... ...")

	ctx1, fn1 := context.WithTimeout(context.Background(), 10*time.Second)
	defer fn1()
	if err := grpc.Shutdown(ctx1); err != nil {
		fmt.Println("grpc shutdown error :", err)
	}

	ctx2, fn2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer fn2()
	if err := ws.Shutdown(ctx2); err != nil {
		fmt.Println("websocket shutdown error :", err)
	}

	ctx3, fn3 := context.WithTimeout(context.Background(), 10*time.Second)
	defer fn3()
	if err := http.Shutdown(ctx3); err != nil {
		fmt.Println("http shutdown error :", err)
	}

	fmt.Println("exit ... ... ...")
}
