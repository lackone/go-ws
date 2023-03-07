package server

import (
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/etcd"
	"github.com/lackone/go-ws/pkg/proto/im"
	"github.com/lackone/go-ws/pkg/service"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func InitGRPCServer() {
	if global.IsCluster() {
		listen, err := net.Listen("tcp", fmt.Sprintf(":%d", global.GrpcSetting.GrpcPort))
		if err != nil {
			fmt.Println("grpc listen error : ", err.Error())
			return
		}
		server := grpc.NewServer()

		im.RegisterIMServiceServer(server, &service.IMService{})

		go func() {
			err = server.Serve(listen)
			if err != nil {
				fmt.Println("grpc serve error : ", err.Error())
				return
			}
		}()

		fmt.Printf("grpc run [:%d] success \n", global.GrpcSetting.GrpcPort)

		//监控信号，实现优雅关机
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("grpc server shutdown ... ...")

		server.GracefulStop()

		fmt.Println("grpc server exit ... ...")
	}
}

var WsEtcdReg *etcd.EtcdRegister
var WsEtcdDis *etcd.EtcdDiscover

// 在ETCD注册服务
func RegisterGRPC() {
	if global.IsCluster() {
		var err error
		WsEtcdReg, err = etcd.NewEtcdRegister(global.EtcdClient, 10)
		if err != nil {
			panic(err)
		}
		addr := net.JoinHostPort(global.LocalIP, strconv.Itoa(global.GrpcSetting.GrpcPort))
		err = WsEtcdReg.RegService(global.ETCD_WS_SERVERS+addr, addr)
		if err != nil {
			panic(err)
		}
		WsEtcdDis, err = etcd.NewEtcdDiscover(global.EtcdClient, global.ETCD_WS_SERVERS)
		if err != nil {
			panic(err)
		}
		fmt.Println(WsEtcdDis.ServiceList())
	}
}
