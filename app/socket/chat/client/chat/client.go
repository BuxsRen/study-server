// 聊天服务器客户端
package chat

import (
	"context"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"study-server/app/socket/chat/chat"
	"study-server/app/socket/chat/client/gateway"
	"study-server/app/socket/chat/client/match"
	"study-server/bootstrap/config"
	gc "study-server/bootstrap/grpc/client"
	pm "study-server/bootstrap/grpc/proto/manage"
	ws "study-server/bootstrap/websocket/client"
)

type Server struct {
	gc pm.ManageGrpcClient
}

// 将聊天服务器注册到管理服务端
func Run() {
	token := encry.EncryptToken(chat.DeviceNumber, -1, map[string]interface{}{
		"group": config.App.Server.Mode,                                  // 注册的组
		"rpc":   config.App.Server.Ip + ":" + config.App.Server.GrpcPort, // 注册的grpc地址
	})

	s := &Server{}
	// 注册grpc到管理服务端
	conn := gc.Dial(config.App.Server.ManageRpc)
	s.gc = pm.NewManageGrpcClient(conn)

	// 将聊天服务器ws注册到管理服务端
	server := ws.New(chat.DeviceNumber, config.App.Server.ManageWs+"?token="+token)
	server.Bind(s)
	server.Start(true, 3)
}

func (s *Server) OnOpen(socket *ws.WebSocket) {
	utils.Println("聊天服务注册成功 <=> 管理服务端")
	matchNode, _ := s.gc.GetService(context.Background(), &pm.Node{Mode: "match"})
	// 连接匹配服务器
	if matchNode != nil {
		for _, v := range matchNode.List {
			go (&match.Match{}).Run(v.Rpc, v.Id)
		}
	}

	gatewayNode, _ := s.gc.GetService(context.Background(), &pm.Node{Mode: "gateway"})
	// 连接匹配服务器
	if gatewayNode != nil {
		for _, v := range gatewayNode.List {
			go (&gateway.Gateway{}).Run(v.Rpc, v.Id)
		}
	}
}

func (s *Server) OnClose(socket *ws.WebSocket) {
	utils.Println("服务连接断开 <=x=> 管理服务端")
}

func (s *Server) OnError(socket *ws.WebSocket) {
	utils.Println("聊天服务注册失败 =>x 管理服务端")
}
