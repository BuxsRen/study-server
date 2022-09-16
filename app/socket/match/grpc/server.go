package grpc

import (
	pm "study-server/bootstrap/grpc/proto/match"
	gs "study-server/bootstrap/grpc/server"
)

// 匹配服务器GRPC服务端
type Server struct {
	pm.UnimplementedMatchGrpcServer
}

// 匹配服务器grpc服务
func Star() {
	server := gs.New()
	pm.RegisterMatchGrpcServer(server.RpcServer, new(Server))
	server.Start()
}
