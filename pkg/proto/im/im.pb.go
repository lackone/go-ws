// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0--rc1
// source: im.proto

package im

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendClientsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Tos  []string `protobuf:"bytes,2,rep,name=tos,proto3" json:"tos,omitempty"`
	Msg  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendClientsReq) Reset() {
	*x = SendClientsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendClientsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendClientsReq) ProtoMessage() {}

func (x *SendClientsReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendClientsReq.ProtoReflect.Descriptor instead.
func (*SendClientsReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0}
}

func (x *SendClientsReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendClientsReq) GetTos() []string {
	if x != nil {
		return x.Tos
	}
	return nil
}

func (x *SendClientsReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SendGroupsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Groups []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Msg    string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendGroupsReq) Reset() {
	*x = SendGroupsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendGroupsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendGroupsReq) ProtoMessage() {}

func (x *SendGroupsReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendGroupsReq.ProtoReflect.Descriptor instead.
func (*SendGroupsReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{1}
}

func (x *SendGroupsReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendGroupsReq) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *SendGroupsReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SendMachinesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Ips  []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	Msg  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMachinesReq) Reset() {
	*x = SendMachinesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMachinesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMachinesReq) ProtoMessage() {}

func (x *SendMachinesReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMachinesReq.ProtoReflect.Descriptor instead.
func (*SendMachinesReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{2}
}

func (x *SendMachinesReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendMachinesReq) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *SendMachinesReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *BroadcastReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AddGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Groups   []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *AddGroupReq) Reset() {
	*x = AddGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupReq) ProtoMessage() {}

func (x *AddGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupReq.ProtoReflect.Descriptor instead.
func (*AddGroupReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{4}
}

func (x *AddGroupReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AddGroupReq) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type DelGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Groups   []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *DelGroupReq) Reset() {
	*x = DelGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelGroupReq) ProtoMessage() {}

func (x *DelGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelGroupReq.ProtoReflect.Descriptor instead.
func (*DelGroupReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{5}
}

func (x *DelGroupReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DelGroupReq) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type OnlineListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnlineListReq) Reset() {
	*x = OnlineListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineListReq) ProtoMessage() {}

func (x *OnlineListReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineListReq.ProtoReflect.Descriptor instead.
func (*OnlineListReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{6}
}

type GroupListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *GroupListReq) Reset() {
	*x = GroupListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupListReq) ProtoMessage() {}

func (x *GroupListReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupListReq.ProtoReflect.Descriptor instead.
func (*GroupListReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{7}
}

func (x *GroupListReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type MachineListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MachineListReq) Reset() {
	*x = MachineListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListReq) ProtoMessage() {}

func (x *MachineListReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListReq.ProtoReflect.Descriptor instead.
func (*MachineListReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{8}
}

type CommonRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CommonRes) Reset() {
	*x = CommonRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRes) ProtoMessage() {}

func (x *CommonRes) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRes.ProtoReflect.Descriptor instead.
func (*CommonRes) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{9}
}

func (x *CommonRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonRes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_im_proto protoreflect.FileDescriptor

var file_im_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x6f, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x4d, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x49, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x34,
	0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x41, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x41, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x2a, 0x0a, 0x0c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x45, 0x0a, 0x09, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x93, 0x03, 0x0a, 0x09, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0a,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a,
	0x53, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c,
	0x2e, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x44, 0x65,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0x2e, 0x44, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_proto_rawDescOnce sync.Once
	file_im_proto_rawDescData = file_im_proto_rawDesc
)

func file_im_proto_rawDescGZIP() []byte {
	file_im_proto_rawDescOnce.Do(func() {
		file_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_proto_rawDescData)
	})
	return file_im_proto_rawDescData
}

var file_im_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_im_proto_goTypes = []interface{}{
	(*SendClientsReq)(nil),  // 0: SendClientsReq
	(*SendGroupsReq)(nil),   // 1: SendGroupsReq
	(*SendMachinesReq)(nil), // 2: SendMachinesReq
	(*BroadcastReq)(nil),    // 3: BroadcastReq
	(*AddGroupReq)(nil),     // 4: AddGroupReq
	(*DelGroupReq)(nil),     // 5: DelGroupReq
	(*OnlineListReq)(nil),   // 6: OnlineListReq
	(*GroupListReq)(nil),    // 7: GroupListReq
	(*MachineListReq)(nil),  // 8: MachineListReq
	(*CommonRes)(nil),       // 9: CommonRes
}
var file_im_proto_depIdxs = []int32{
	0, // 0: IMService.SendClients:input_type -> SendClientsReq
	1, // 1: IMService.SendGroups:input_type -> SendGroupsReq
	2, // 2: IMService.SendMachines:input_type -> SendMachinesReq
	3, // 3: IMService.Broadcast:input_type -> BroadcastReq
	4, // 4: IMService.AddGroup:input_type -> AddGroupReq
	5, // 5: IMService.DelGroup:input_type -> DelGroupReq
	6, // 6: IMService.OnlineList:input_type -> OnlineListReq
	7, // 7: IMService.GroupList:input_type -> GroupListReq
	8, // 8: IMService.MachineList:input_type -> MachineListReq
	9, // 9: IMService.SendClients:output_type -> CommonRes
	9, // 10: IMService.SendGroups:output_type -> CommonRes
	9, // 11: IMService.SendMachines:output_type -> CommonRes
	9, // 12: IMService.Broadcast:output_type -> CommonRes
	9, // 13: IMService.AddGroup:output_type -> CommonRes
	9, // 14: IMService.DelGroup:output_type -> CommonRes
	9, // 15: IMService.OnlineList:output_type -> CommonRes
	9, // 16: IMService.GroupList:output_type -> CommonRes
	9, // 17: IMService.MachineList:output_type -> CommonRes
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_proto_init() }
func file_im_proto_init() {
	if File_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendClientsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendGroupsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMachinesReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelGroupReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_im_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_proto_goTypes,
		DependencyIndexes: file_im_proto_depIdxs,
		MessageInfos:      file_im_proto_msgTypes,
	}.Build()
	File_im_proto = out.File
	file_im_proto_rawDesc = nil
	file_im_proto_goTypes = nil
	file_im_proto_depIdxs = nil
}
