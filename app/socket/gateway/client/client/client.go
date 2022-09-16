package client

import (
	"context"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"study-server/app/socket/gateway/client/chat"
	"study-server/app/socket/gateway/client/game"
	"study-server/app/socket/gateway/client/match"
	"study-server/app/socket/gateway/gateway"
	"study-server/bootstrap/config"
	gc "study-server/bootstrap/grpc/client"
	pm "study-server/bootstrap/grpc/proto/manage"
	ws "study-server/bootstrap/websocket/client"
)

type Server struct {
	gc pm.ManageGrpcClient
}

// 将网关注册到管理服务端
func Run() {
	token := encry.EncryptToken(gateway.DeviceNumber, -1, map[string]interface{}{
		"group": config.App.Server.Mode,                                  // 注册的组
		"rpc":   config.App.Server.Ip + ":" + config.App.Server.GrpcPort, // 注册的grpc地址
	})
	// 将网关ws注册到管理服务端
	s := &Server{}
	// 注册grpc到管理服务端
	conn := gc.Dial(config.App.Server.ManageRpc)
	s.gc = pm.NewManageGrpcClient(conn)

	server := ws.New(gateway.DeviceNumber, config.App.Server.ManageWs+"?token="+token)
	server.Bind(s)
	server.Start(true, 3)
}

func (s *Server) OnOpen(socket *ws.WebSocket) {
	utils.Println("网关服务 <==> 注册成功 <=> 管理服务器")
	// 获取聊天服务器列表
	chatNode, _ := s.gc.GetService(context.Background(), &pm.Node{Mode: "chat"})

	// 连接聊天服务器
	if chatNode != nil {
		for _, v := range chatNode.List {
			c := &chat.Chat{}
			go c.Run(v.Rpc, v.Id)
		}
	}

	matchNode, _ := s.gc.GetService(context.Background(), &pm.Node{Mode: "match"})

	// 连接匹配服务器
	if matchNode != nil {
		for _, v := range matchNode.List {
			go (&match.Match{}).Run(v.Rpc, v.Id)
		}
	}

	gameNode, _ := s.gc.GetService(context.Background(), &pm.Node{Mode: "game"})

	// 连接匹配服务器
	if gameNode != nil {
		for _, v := range gameNode.List {
			go (&game.Game{}).Run(v.Rpc, v.Id, v.Udp)
		}
	}
}

func (s *Server) OnClose(socket *ws.WebSocket) {
	utils.Println("服务连接断开 <=x=> 管理服务器")
}

func (s *Server) OnError(socket *ws.WebSocket) {
	utils.Println("网关服务注册失败 =>x 管理服务器")
}