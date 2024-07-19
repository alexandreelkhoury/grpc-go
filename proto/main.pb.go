// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/main.proto

package grpc_go

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Reputation int32  `protobuf:"varint,2,opt,name=reputation,proto3" json:"reputation,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RegisterResponse) GetReputation() int32 {
	if x != nil {
		return x.Reputation
	}
	return 0
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash         string `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	PreviousBlockHash string `protobuf:"bytes,2,opt,name=previous_block_hash,json=previousBlockHash,proto3" json:"previous_block_hash,omitempty"`
	BlockNumber       int32  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Data              string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *BlockInfo) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockInfo) GetPreviousBlockHash() string {
	if x != nil {
		return x.PreviousBlockHash
	}
	return ""
}

func (x *BlockInfo) GetBlockNumber() int32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockInfo) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount   int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Data     string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Transaction) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Transaction) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type BakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *BakeRequest) Reset() {
	*x = BakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeRequest) ProtoMessage() {}

func (x *BakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeRequest.ProtoReflect.Descriptor instead.
func (*BakeRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *BakeRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type BakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BakeResponse) Reset() {
	*x = BakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeResponse) ProtoMessage() {}

func (x *BakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeResponse.ProtoReflect.Descriptor instead.
func (*BakeResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{7}
}

func (x *BakeResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BakeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ConfirmRequest) Reset() {
	*x = ConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRequest) ProtoMessage() {}

func (x *ConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRequest.ProtoReflect.Descriptor instead.
func (*ConfirmRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{8}
}

func (x *ConfirmRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x91, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x6d, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x21, 0x0a, 0x0b, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x81, 0x04, 0x0a, 0x0a, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x4f, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x09, 0x42, 0x61, 0x6b, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x42, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x61, 0x6b, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65,
	0x78, 0x61, 0x6e, 0x64, 0x72, 0x65, 0x65, 0x6c, 0x6b, 0x68, 0x6f, 0x75, 0x72, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_main_proto_goTypes = []any{
	(*Empty)(nil),             // 0: grpc_golang_template.Empty
	(*RegisterResponse)(nil),  // 1: grpc_golang_template.RegisterResponse
	(*SubscribeRequest)(nil),  // 2: grpc_golang_template.SubscribeRequest
	(*SubscribeResponse)(nil), // 3: grpc_golang_template.SubscribeResponse
	(*BlockInfo)(nil),         // 4: grpc_golang_template.BlockInfo
	(*Transaction)(nil),       // 5: grpc_golang_template.Transaction
	(*BakeRequest)(nil),       // 6: grpc_golang_template.BakeRequest
	(*BakeResponse)(nil),      // 7: grpc_golang_template.BakeResponse
	(*ConfirmRequest)(nil),    // 8: grpc_golang_template.ConfirmRequest
}
var file_proto_main_proto_depIdxs = []int32{
	0, // 0: grpc_golang_template.Blockchain.Register:input_type -> grpc_golang_template.Empty
	2, // 1: grpc_golang_template.Blockchain.Subscribe:input_type -> grpc_golang_template.SubscribeRequest
	0, // 2: grpc_golang_template.Blockchain.GetLastBlock:input_type -> grpc_golang_template.Empty
	5, // 3: grpc_golang_template.Blockchain.AddTransaction:input_type -> grpc_golang_template.Transaction
	6, // 4: grpc_golang_template.Blockchain.BakeBlock:input_type -> grpc_golang_template.BakeRequest
	8, // 5: grpc_golang_template.Blockchain.ConfirmBake:input_type -> grpc_golang_template.ConfirmRequest
	1, // 6: grpc_golang_template.Blockchain.Register:output_type -> grpc_golang_template.RegisterResponse
	3, // 7: grpc_golang_template.Blockchain.Subscribe:output_type -> grpc_golang_template.SubscribeResponse
	4, // 8: grpc_golang_template.Blockchain.GetLastBlock:output_type -> grpc_golang_template.BlockInfo
	0, // 9: grpc_golang_template.Blockchain.AddTransaction:output_type -> grpc_golang_template.Empty
	7, // 10: grpc_golang_template.Blockchain.BakeBlock:output_type -> grpc_golang_template.BakeResponse
	0, // 11: grpc_golang_template.Blockchain.ConfirmBake:output_type -> grpc_golang_template.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_main_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_proto_main_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterResponse); i {
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
		file_proto_main_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SubscribeRequest); i {
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
		file_proto_main_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SubscribeResponse); i {
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
		file_proto_main_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BlockInfo); i {
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
		file_proto_main_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Transaction); i {
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
		file_proto_main_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*BakeRequest); i {
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
		file_proto_main_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*BakeResponse); i {
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
		file_proto_main_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmRequest); i {
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
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}
