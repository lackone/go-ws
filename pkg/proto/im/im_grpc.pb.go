// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc1
// source: im.proto

package im

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IMServiceClient is the client API for IMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IMServiceClient interface {
	SendClients(ctx context.Context, in *SendClientsReq, opts ...grpc.CallOption) (*CommonRes, error)
	SendGroups(ctx context.Context, in *SendGroupsReq, opts ...grpc.CallOption) (*CommonRes, error)
	SendMachines(ctx context.Context, in *SendMachinesReq, opts ...grpc.CallOption) (*CommonRes, error)
	Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*CommonRes, error)
	AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*CommonRes, error)
	DelGroup(ctx context.Context, in *DelGroupReq, opts ...grpc.CallOption) (*CommonRes, error)
	OnlineList(ctx context.Context, in *OnlineListReq, opts ...grpc.CallOption) (*CommonRes, error)
	GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*CommonRes, error)
	MachineList(ctx context.Context, in *MachineListReq, opts ...grpc.CallOption) (*CommonRes, error)
}

type iMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIMServiceClient(cc grpc.ClientConnInterface) IMServiceClient {
	return &iMServiceClient{cc}
}

func (c *iMServiceClient) SendClients(ctx context.Context, in *SendClientsReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/SendClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) SendGroups(ctx context.Context, in *SendGroupsReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/SendGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) SendMachines(ctx context.Context, in *SendMachinesReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/SendMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) DelGroup(ctx context.Context, in *DelGroupReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/DelGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) OnlineList(ctx context.Context, in *OnlineListReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/OnlineList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/GroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMServiceClient) MachineList(ctx context.Context, in *MachineListReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/IMService/MachineList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMServiceServer is the server API for IMService service.
// All implementations must embed UnimplementedIMServiceServer
// for forward compatibility
type IMServiceServer interface {
	SendClients(context.Context, *SendClientsReq) (*CommonRes, error)
	SendGroups(context.Context, *SendGroupsReq) (*CommonRes, error)
	SendMachines(context.Context, *SendMachinesReq) (*CommonRes, error)
	Broadcast(context.Context, *BroadcastReq) (*CommonRes, error)
	AddGroup(context.Context, *AddGroupReq) (*CommonRes, error)
	DelGroup(context.Context, *DelGroupReq) (*CommonRes, error)
	OnlineList(context.Context, *OnlineListReq) (*CommonRes, error)
	GroupList(context.Context, *GroupListReq) (*CommonRes, error)
	MachineList(context.Context, *MachineListReq) (*CommonRes, error)
	mustEmbedUnimplementedIMServiceServer()
}

// UnimplementedIMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIMServiceServer struct {
}

func (UnimplementedIMServiceServer) SendClients(context.Context, *SendClientsReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClients not implemented")
}
func (UnimplementedIMServiceServer) SendGroups(context.Context, *SendGroupsReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroups not implemented")
}
func (UnimplementedIMServiceServer) SendMachines(context.Context, *SendMachinesReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMachines not implemented")
}
func (UnimplementedIMServiceServer) Broadcast(context.Context, *BroadcastReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedIMServiceServer) AddGroup(context.Context, *AddGroupReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedIMServiceServer) DelGroup(context.Context, *DelGroupReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelGroup not implemented")
}
func (UnimplementedIMServiceServer) OnlineList(context.Context, *OnlineListReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineList not implemented")
}
func (UnimplementedIMServiceServer) GroupList(context.Context, *GroupListReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedIMServiceServer) MachineList(context.Context, *MachineListReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MachineList not implemented")
}
func (UnimplementedIMServiceServer) mustEmbedUnimplementedIMServiceServer() {}

// UnsafeIMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IMServiceServer will
// result in compilation errors.
type UnsafeIMServiceServer interface {
	mustEmbedUnimplementedIMServiceServer()
}

func RegisterIMServiceServer(s grpc.ServiceRegistrar, srv IMServiceServer) {
	s.RegisterService(&IMService_ServiceDesc, srv)
}

func _IMService_SendClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendClientsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).SendClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/SendClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).SendClients(ctx, req.(*SendClientsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_SendGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).SendGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/SendGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).SendGroups(ctx, req.(*SendGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_SendMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMachinesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).SendMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/SendMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).SendMachines(ctx, req.(*SendMachinesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).Broadcast(ctx, req.(*BroadcastReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).AddGroup(ctx, req.(*AddGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_DelGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).DelGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/DelGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).DelGroup(ctx, req.(*DelGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_OnlineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).OnlineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/OnlineList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).OnlineList(ctx, req.(*OnlineListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/GroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).GroupList(ctx, req.(*GroupListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IMService_MachineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServiceServer).MachineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IMService/MachineList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServiceServer).MachineList(ctx, req.(*MachineListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IMService_ServiceDesc is the grpc.ServiceDesc for IMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IMService",
	HandlerType: (*IMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClients",
			Handler:    _IMService_SendClients_Handler,
		},
		{
			MethodName: "SendGroups",
			Handler:    _IMService_SendGroups_Handler,
		},
		{
			MethodName: "SendMachines",
			Handler:    _IMService_SendMachines_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _IMService_Broadcast_Handler,
		},
		{
			MethodName: "AddGroup",
			Handler:    _IMService_AddGroup_Handler,
		},
		{
			MethodName: "DelGroup",
			Handler:    _IMService_DelGroup_Handler,
		},
		{
			MethodName: "OnlineList",
			Handler:    _IMService_OnlineList_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _IMService_GroupList_Handler,
		},
		{
			MethodName: "MachineList",
			Handler:    _IMService_MachineList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im.proto",
}