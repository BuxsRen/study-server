package grpc

import (
	pc "study-server/bootstrap/grpc/proto/chat"
	gs "study-server/bootstrap/grpc/server"
)

// 管理服务
type Server struct {
	pc.UnimplementedChatGrpcServer
}

// 聊天服务端grpc服务
func Star() {
	server := gs.New()
	pc.RegisterChatGrpcServer(server.RpcServer, new(Server))
	server.Start()
}