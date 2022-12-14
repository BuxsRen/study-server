// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: gateway.proto

package gateway

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

// GatewayGrpcClient is the client API for GatewayGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayGrpcClient interface {
	// 其他网关用户登录消息监听
	OnLogin(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error)
	// 其他网关用户登录消息监听
	OnLogout(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error)
	// 其他网关用户登录消息监听
	OnAreaJoin(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error)
	// 其他网关用户登录消息监听
	OnAreaExit(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error)
	// 其他网关用户登录消息监听
	OnSendMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error)
}

type gatewayGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayGrpcClient(cc grpc.ClientConnInterface) GatewayGrpcClient {
	return &gatewayGrpcClient{cc}
}

func (c *gatewayGrpcClient) OnLogin(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gateway.GatewayGrpc/OnLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayGrpcClient) OnLogout(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gateway.GatewayGrpc/OnLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayGrpcClient) OnAreaJoin(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gateway.GatewayGrpc/OnAreaJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayGrpcClient) OnAreaExit(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gateway.GatewayGrpc/OnAreaExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayGrpcClient) OnSendMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/gateway.GatewayGrpc/OnSendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayGrpcServer is the server API for GatewayGrpc service.
// All implementations must embed UnimplementedGatewayGrpcServer
// for forward compatibility
type GatewayGrpcServer interface {
	// 其他网关用户登录消息监听
	OnLogin(context.Context, *Msg) (*Res, error)
	// 其他网关用户登录消息监听
	OnLogout(context.Context, *Msg) (*Res, error)
	// 其他网关用户登录消息监听
	OnAreaJoin(context.Context, *Msg) (*Res, error)
	// 其他网关用户登录消息监听
	OnAreaExit(context.Context, *Msg) (*Res, error)
	// 其他网关用户登录消息监听
	OnSendMsg(context.Context, *Msg) (*Res, error)
	mustEmbedUnimplementedGatewayGrpcServer()
}

// UnimplementedGatewayGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayGrpcServer struct {
}

func (UnimplementedGatewayGrpcServer) OnLogin(context.Context, *Msg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnLogin not implemented")
}
func (UnimplementedGatewayGrpcServer) OnLogout(context.Context, *Msg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnLogout not implemented")
}
func (UnimplementedGatewayGrpcServer) OnAreaJoin(context.Context, *Msg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnAreaJoin not implemented")
}
func (UnimplementedGatewayGrpcServer) OnAreaExit(context.Context, *Msg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnAreaExit not implemented")
}
func (UnimplementedGatewayGrpcServer) OnSendMsg(context.Context, *Msg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSendMsg not implemented")
}
func (UnimplementedGatewayGrpcServer) mustEmbedUnimplementedGatewayGrpcServer() {}

// UnsafeGatewayGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayGrpcServer will
// result in compilation errors.
type UnsafeGatewayGrpcServer interface {
	mustEmbedUnimplementedGatewayGrpcServer()
}

func RegisterGatewayGrpcServer(s grpc.ServiceRegistrar, srv GatewayGrpcServer) {
	s.RegisterService(&GatewayGrpc_ServiceDesc, srv)
}

func _GatewayGrpc_OnLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayGrpcServer).OnLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayGrpc/OnLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayGrpcServer).OnLogin(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayGrpc_OnLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayGrpcServer).OnLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayGrpc/OnLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayGrpcServer).OnLogout(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayGrpc_OnAreaJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayGrpcServer).OnAreaJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayGrpc/OnAreaJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayGrpcServer).OnAreaJoin(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayGrpc_OnAreaExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayGrpcServer).OnAreaExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayGrpc/OnAreaExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayGrpcServer).OnAreaExit(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayGrpc_OnSendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayGrpcServer).OnSendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayGrpc/OnSendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayGrpcServer).OnSendMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayGrpc_ServiceDesc is the grpc.ServiceDesc for GatewayGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GatewayGrpc",
	HandlerType: (*GatewayGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnLogin",
			Handler:    _GatewayGrpc_OnLogin_Handler,
		},
		{
			MethodName: "OnLogout",
			Handler:    _GatewayGrpc_OnLogout_Handler,
		},
		{
			MethodName: "OnAreaJoin",
			Handler:    _GatewayGrpc_OnAreaJoin_Handler,
		},
		{
			MethodName: "OnAreaExit",
			Handler:    _GatewayGrpc_OnAreaExit_Handler,
		},
		{
			MethodName: "OnSendMsg",
			Handler:    _GatewayGrpc_OnSendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
