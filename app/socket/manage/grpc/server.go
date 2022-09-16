package grpc

import (
	pm "study-server/bootstrap/grpc/proto/manage"
	gs "study-server/bootstrap/grpc/server"
)

// 管理服务
type Server struct {
	pm.UnimplementedManageGrpcServer
}

// 管理服务端grpc服务
func Star() {
	server := gs.New()
	pm.RegisterManageGrpcServer(server.RpcServer, new(Server))
	server.Start()
}