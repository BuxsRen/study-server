package client

import (
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"study-server/app/socket/game/game"
	"study-server/bootstrap/config"
	ws "study-server/bootstrap/websocket/client"
)

type Server struct{}

// 游戏服务器注册到管理服务器
func Run() {
	token := encry.EncryptToken(game.DeviceNumber, -1, map[string]interface{}{
		"group": config.App.Server.Mode,                                  // 注册的组
		"rpc":   config.App.Server.Ip + ":" + config.App.Server.GrpcPort, // 注册的grpc地址
		"udp":   config.App.Server.Ip + ":" + config.App.Server.Port,
	})
	s := &Server{}
	// 将聊天服务器ws注册到管理服务端
	server := ws.New(game.DeviceNumber, config.App.Server.ManageWs+"?token="+token)
	server.Bind(s)
	server.Start(true, 3)
}

func (s *Server) OnOpen(socket *ws.WebSocket) {
	utils.Println("游戏服务注册成功 <=> 管理服务端")
}

func (s *Server) OnClose(socket *ws.WebSocket) {
	utils.Println("游戏连接断开 <=x=> 管理服务端")
}

func (s *Server) OnError(socket *ws.WebSocket) {
	utils.Println("游戏服务注册失败 =>x 管理服务端")
}
