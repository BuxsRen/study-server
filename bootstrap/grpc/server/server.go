package server

import (
	"fmt"
	"google.golang.org/grpc"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
	"net"
)

type GrpcServer struct {
	RpcServer *grpc.Server
	lis       net.Listener
	err       error
}

// 初始化，创建rpc监听服务
func New() *GrpcServer {
	g := &GrpcServer{}
	g.RpcServer = grpc.NewServer()
	g.lis, g.err = net.Listen("tcp", fmt.Sprintf("%v:%v", config.App.Server.Host, config.App.Server.GrpcPort))
	if g.err != nil {
		fmt.Println("➦ " + g.err.Error())
		(&helper.Helper{}).Exit("✘ This Port Is Already In Use !", 3)
	}
	ip := "127.0.0.1"
	if config.App.Server.Host != "127.0.0.1" {
		ip = config.App.Server.Ip
	} else if config.App.Server.Host != "0.0.0.0" {
		ip = config.App.Server.Host
	}
	fmt.Printf("📡 rpc -> %v:%v\n", ip, config.App.Server.GrpcPort)
	fmt.Printf("► OK! Start Service...\n\n")
	return g
}

// 开始 GRPC 服务端
func (g *GrpcServer) Start() {
	g.err = g.RpcServer.Serve(g.lis)
	if g.err != nil {
		(&helper.Helper{}).Exit("✘ "+g.err.Error(), 3)
	}
}