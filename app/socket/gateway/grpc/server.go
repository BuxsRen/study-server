package grpc

import (
	pg "study-server/bootstrap/grpc/proto/gateway"
	gs "study-server/bootstrap/grpc/server"
)

// 管理服务
type Server struct {
	pg.UnimplementedGatewayGrpcServer
}

// 网关grpc服务
func Star() {
	server := gs.New()
	pg.RegisterGatewayGrpcServer(server.RpcServer, new(Server))
	server.Start()
}
