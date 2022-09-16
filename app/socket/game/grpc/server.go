package grpc

import (
	pg "study-server/bootstrap/grpc/proto/game"
	gs "study-server/bootstrap/grpc/server"
)

// 管理服务
type Server struct {
	pg.UnimplementedGameGrpcServer
}

// 聊天服务端grpc服务
func Star() {
	server := gs.New()
	pg.RegisterGameGrpcServer(server.RpcServer, new(Server))
	server.Start()
}