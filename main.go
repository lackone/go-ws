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
}

func main() {
	go server.InitHttpServer()

	go server.InitWsServer()

	select {}
}
