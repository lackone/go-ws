package main

import (
	"flag"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/server"
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
	go server.InitGRPCServer()

	//注册服务
	server.RegisterGRPC()

	//启动ws
	go server.InitWsServer()

	//启动http
	server.InitHttpServer()
}
