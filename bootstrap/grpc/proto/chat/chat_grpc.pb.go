// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: chat.proto

package chat

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

// ChatGrpcClient is the client API for ChatGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatGrpcClient interface {
	// 客户端登录通知
	Login(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	// 客户端退出通知
	Logout(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	// 加入区域消息转发
	JoinArea(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	// 退出区域消息转发
	ExitArea(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	// 给用户所在区域广播消息
	SendMsgArea(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*Res, error)
	// 给指定用户发送消息
	SendMsg(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*Res, error)
}

type chatGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewChatGrpcClient(cc grpc.ClientConnInterface) ChatGrpcClient {
	return &chatGrpcClient{cc}
}

func (c *chatGrpcClient) Login(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGrpcClient) Logout(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGrpcClient) JoinArea(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/JoinArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGrpcClient) ExitArea(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/ExitArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGrpcClient) SendMsgArea(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/SendMsgArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGrpcClient) SendMsg(ctx context.Context, in *ReqMsg, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chat.ChatGrpc/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatGrpcServer is the server API for ChatGrpc service.
// All implementations must embed UnimplementedChatGrpcServer
// for forward compatibility
type ChatGrpcServer interface {
	// 客户端登录通知
	Login(context.Context, *Req) (*Res, error)
	// 客户端退出通知
	Logout(context.Context, *Req) (*Res, error)
	// 加入区域消息转发
	JoinArea(context.Context, *Req) (*Res, error)
	// 退出区域消息转发
	ExitArea(context.Context, *Req) (*Res, error)
	// 给用户所在区域广播消息
	SendMsgArea(context.Context, *ReqMsg) (*Res, error)
	// 给指定用户发送消息
	SendMsg(context.Context, *ReqMsg) (*Res, error)
	mustEmbedUnimplementedChatGrpcServer()
}

// UnimplementedChatGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedChatGrpcServer struct {
}

func (UnimplementedChatGrpcServer) Login(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatGrpcServer) Logout(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedChatGrpcServer) JoinArea(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinArea not implemented")
}
func (UnimplementedChatGrpcServer) ExitArea(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitArea not implemented")
}
func (UnimplementedChatGrpcServer) SendMsgArea(context.Context, *ReqMsg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgArea not implemented")
}
func (UnimplementedChatGrpcServer) SendMsg(context.Context, *ReqMsg) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatGrpcServer) mustEmbedUnimplementedChatGrpcServer() {}

// UnsafeChatGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatGrpcServer will
// result in compilation errors.
type UnsafeChatGrpcServer interface {
	mustEmbedUnimplementedChatGrpcServer()
}

func RegisterChatGrpcServer(s grpc.ServiceRegistrar, srv ChatGrpcServer) {
	s.RegisterService(&ChatGrpc_ServiceDesc, srv)
}

func _ChatGrpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).Login(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGrpc_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).Logout(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGrpc_JoinArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).JoinArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/JoinArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).JoinArea(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGrpc_ExitArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).ExitArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/ExitArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).ExitArea(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGrpc_SendMsgArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).SendMsgArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/SendMsgArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).SendMsgArea(ctx, req.(*ReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGrpc_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGrpcServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatGrpc/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGrpcServer).SendMsg(ctx, req.(*ReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatGrpc_ServiceDesc is the grpc.ServiceDesc for ChatGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatGrpc",
	HandlerType: (*ChatGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatGrpc_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ChatGrpc_Logout_Handler,
		},
		{
			MethodName: "JoinArea",
			Handler:    _ChatGrpc_JoinArea_Handler,
		},
		{
			MethodName: "ExitArea",
			Handler:    _ChatGrpc_ExitArea_Handler,
		},
		{
			MethodName: "SendMsgArea",
			Handler:    _ChatGrpc_SendMsgArea_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _ChatGrpc_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
