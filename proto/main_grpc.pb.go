// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: proto/main.proto

package grpc_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Blockchain_Register_FullMethodName       = "/grpc_golang_template.Blockchain/Register"
	Blockchain_Subscribe_FullMethodName      = "/grpc_golang_template.Blockchain/Subscribe"
	Blockchain_GetLastBlock_FullMethodName   = "/grpc_golang_template.Blockchain/GetLastBlock"
	Blockchain_AddTransaction_FullMethodName = "/grpc_golang_template.Blockchain/AddTransaction"
	Blockchain_BakeBlock_FullMethodName      = "/grpc_golang_template.Blockchain/BakeBlock"
	Blockchain_ConfirmBake_FullMethodName    = "/grpc_golang_template.Blockchain/ConfirmBake"
)

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainClient interface {
	Register(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegisterResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	GetLastBlock(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockInfo, error)
	AddTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error)
	BakeBlock(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error)
	ConfirmBake(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Empty, error)
}

type blockchainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainClient(cc grpc.ClientConnInterface) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) Register(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Blockchain_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, Blockchain_Subscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetLastBlock(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockInfo)
	err := c.cc.Invoke(ctx, Blockchain_GetLastBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) AddTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Blockchain_AddTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) BakeBlock(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BakeResponse)
	err := c.cc.Invoke(ctx, Blockchain_BakeBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) ConfirmBake(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Blockchain_ConfirmBake_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
// All implementations must embed UnimplementedBlockchainServer
// for forward compatibility
type BlockchainServer interface {
	Register(context.Context, *Empty) (*RegisterResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	GetLastBlock(context.Context, *Empty) (*BlockInfo, error)
	AddTransaction(context.Context, *Transaction) (*Empty, error)
	BakeBlock(context.Context, *BakeRequest) (*BakeResponse, error)
	ConfirmBake(context.Context, *ConfirmRequest) (*Empty, error)
	mustEmbedUnimplementedBlockchainServer()
}

// UnimplementedBlockchainServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (UnimplementedBlockchainServer) Register(context.Context, *Empty) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBlockchainServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedBlockchainServer) GetLastBlock(context.Context, *Empty) (*BlockInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastBlock not implemented")
}
func (UnimplementedBlockchainServer) AddTransaction(context.Context, *Transaction) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedBlockchainServer) BakeBlock(context.Context, *BakeRequest) (*BakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BakeBlock not implemented")
}
func (UnimplementedBlockchainServer) ConfirmBake(context.Context, *ConfirmRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmBake not implemented")
}
func (UnimplementedBlockchainServer) mustEmbedUnimplementedBlockchainServer() {}

// UnsafeBlockchainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServer will
// result in compilation errors.
type UnsafeBlockchainServer interface {
	mustEmbedUnimplementedBlockchainServer()
}

func RegisterBlockchainServer(s grpc.ServiceRegistrar, srv BlockchainServer) {
	s.RegisterService(&Blockchain_ServiceDesc, srv)
}

func _Blockchain_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).Register(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetLastBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetLastBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_GetLastBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetLastBlock(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_AddTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).AddTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_BakeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).BakeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_BakeBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).BakeBlock(ctx, req.(*BakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_ConfirmBake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).ConfirmBake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_ConfirmBake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).ConfirmBake(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blockchain_ServiceDesc is the grpc.ServiceDesc for Blockchain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blockchain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_golang_template.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Blockchain_Register_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Blockchain_Subscribe_Handler,
		},
		{
			MethodName: "GetLastBlock",
			Handler:    _Blockchain_GetLastBlock_Handler,
		},
		{
			MethodName: "AddTransaction",
			Handler:    _Blockchain_AddTransaction_Handler,
		},
		{
			MethodName: "BakeBlock",
			Handler:    _Blockchain_BakeBlock_Handler,
		},
		{
			MethodName: "ConfirmBake",
			Handler:    _Blockchain_ConfirmBake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}
