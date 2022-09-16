package client

import (
	"encoding/json"
	"study-server/app/libs/utils"
	"study-server/app/socket/gateway/client/chat"
	"study-server/app/socket/gateway/client/game"
	"study-server/app/socket/gateway/client/match"
	pm "study-server/bootstrap/grpc/proto/manage"
	ws "study-server/bootstrap/websocket/client"
)

var _ ws.WebSocketInterface = (*Server)(nil)

type RegisterMsg struct {
	Action string
	Node   pm.Node
}

// 服务注册通知
func (s *Server) ServerRegister(socket *ws.WebSocket, msg []byte) {
	var data RegisterMsg
	_ = json.Unmarshal(msg, &data)
	switch data.Node.Mode {
	case "chat":
		utils.Println("聊天服务器上线")
		c := &chat.Chat{}
		go c.Run(data.Node.Rpc, data.Node.Id)
	case "match":
		utils.Println("匹配服务器上线")
		go (&match.Match{}).Run(data.Node.Rpc, data.Node.Id)
	case "game":
		utils.Println("游戏服务器上线")
		go (&game.Game{}).Run(data.Node.Rpc, data.Node.Id, data.Node.Udp)
	}
}

// 服务注销通知
func (s *Server) ServerLogout(socket *ws.WebSocket, msg []byte) {
	var data RegisterMsg
	_ = json.Unmarshal(msg, &data)
	switch data.Node.Mode {
	case "chat":
		utils.Println("聊天服务器离线")
		chat.ChatGrpc.Delete(data.Node.Id)
	case "match":
		utils.Println("匹配服务器离线")
		match.MatchService = nil
	case "game":
		utils.Println("游戏服务器离线")
		game.GameGrpc.Delete(data.Node.Id)
	}
}
