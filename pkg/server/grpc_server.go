package server

import (
	"context"
	"fmt"
	"github.com/lackone/go-ws/global"
	"github.com/lackone/go-ws/pkg/etcd"
	"github.com/lackone/go-ws/pkg/proto/im"
	"github.com/lackone/go-ws/pkg/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (g *GRPCServer) start() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", global.GrpcSetting.GrpcPort))
	if err != nil {
		panic(err)
	}
	g.server = grpc.NewServer()

	im.RegisterIMServiceServer(g.server, &service.IMService{})

	go func() {
		err = g.server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("grpc[:%d] success \n", global.GrpcSetting.GrpcPort)
}

// 运行
func (g *GRPCServer) Run() {
	if global.IsCluster() {
		g.start()

		g.registerGRPC()
	}
}

// 退出
func (g *GRPCServer) Shutdown(ctx context.Context) error {
	ok := make(chan struct{})

	go func(ok chan struct{}) {
		fmt.Printf("grpc[:%d] shutdown \n", global.GrpcSetting.GrpcPort)

		g.server.GracefulStop()

		g.clearEtcdData()

		fmt.Printf("grpc[:%d] exit \n", global.GrpcSetting.GrpcPort)

		ok <- struct{}{}
	}(ok)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ok:
		return nil
	}
}

// 清除ETCD中的数据
func (g *GRPCServer) clearEtcdData() {
	global.EtcdKV.DelAll(global.ETCD_WS_ACCOUNTS + "/" + global.LocalIP)
	global.EtcdKV.DelAll(global.ETCD_WS_MACHINES + "/" + global.LocalIP)
}

// 在ETCD注册服务
func (g *GRPCServer) registerGRPC() {
	if global.IsCluster() {
		reg, err := etcd.NewEtcdRegister(global.EtcdClient, 10)
		if err != nil {
			panic(err)
		}
		addr := net.JoinHostPort(global.LocalIP, strconv.Itoa(global.GrpcSetting.GrpcPort))
		err = reg.RegService(global.ETCD_WS_SERVICES+addr, addr)
		if err != nil {
			panic(err)
		}
	}
}
