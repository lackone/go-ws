package client

import (
	"context"
	"github.com/lackone/go-ws/pkg/proto/im"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IMGrpcClient struct {
	conn   *grpc.ClientConn
	client im.IMServiceClient
}

func NewIMGrpcClient(addr string) (*IMGrpcClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := im.NewIMServiceClient(conn)

	return &IMGrpcClient{
		conn:   conn,
		client: client,
	}, nil
}

func (i *IMGrpcClient) Close() {
	i.conn.Close()
}

func (i *IMGrpcClient) SendClients(from string, tos []string, msg string) (*im.CommonRes, error) {
	res, err := i.client.SendClients(context.Background(), &im.SendClientsReq{
		From: from,
		Tos:  tos,
		Msg:  msg,
	})

	return res, err
}

func (i *IMGrpcClient) SendGroups(from string, groups []string, msg string) (*im.CommonRes, error) {
	res, err := i.client.SendGroups(context.Background(), &im.SendGroupsReq{
		From:   from,
		Groups: groups,
		Msg:    msg,
	})
	return res, err
}

func (i *IMGrpcClient) SendMachines(from string, ips []string, msg string) (*im.CommonRes, error) {
	res, err := i.client.SendMachines(context.Background(), &im.SendMachinesReq{
		From: from,
		Ips:  ips,
		Msg:  msg,
	})
	return res, err
}

func (i *IMGrpcClient) Broadcast(from string, msg string) (*im.CommonRes, error) {
	res, err := i.client.Broadcast(context.Background(), &im.BroadcastReq{
		From: from,
		Msg:  msg,
	})
	return res, err
}

func (i *IMGrpcClient) AddGroup(clientId string, groups []string) (*im.CommonRes, error) {
	res, err := i.client.AddGroup(context.Background(), &im.AddGroupReq{
		ClientId: clientId,
		Groups:   groups,
	})
	return res, err
}

func (i *IMGrpcClient) DelGroup(clientId string, groups []string) (*im.CommonRes, error) {
	res, err := i.client.DelGroup(context.Background(), &im.DelGroupReq{
		ClientId: clientId,
		Groups:   groups,
	})
	return res, err
}

func (i *IMGrpcClient) OnlineList() (*im.CommonRes, error) {
	res, err := i.client.OnlineList(context.Background(), &im.OnlineListReq{})
	return res, err
}

func (i *IMGrpcClient) GroupList(clientId string) (*im.CommonRes, error) {
	res, err := i.client.GroupList(context.Background(), &im.GroupListReq{
		ClientId: clientId,
	})
	return res, err
}

func (i *IMGrpcClient) MachineList() (*im.CommonRes, error) {
	res, err := i.client.MachineList(context.Background(), &im.MachineListReq{})
	return res, err
}
