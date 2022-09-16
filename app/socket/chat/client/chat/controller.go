package chat

import (
	"encoding/json"
	"study-server/app/libs/utils"
	"study-server/app/socket/chat/client/gateway"
	"study-server/app/socket/chat/client/match"
	pm "study-server/bootstrap/grpc/proto/manage"
	ws "study-server/bootstrap/websocket/client"
)

type RegisterMsg struct {
	Action string
	Node   pm.Node
}

// 服务注册通知
func (s *Server) ServerRegister(socket *ws.WebSocket, msg []byte) {
	var data RegisterMsg
	_ = json.Unmarshal(msg, &data)
	switch data.Node.Mode {
	case "match":
		utils.Println("匹配服务器上线")
		go (&match.Match{}).Run(data.Node.Rpc, data.Node.Id)
	case "gateway":
		utils.Println("网关上线")
		go (&gateway.Gateway{}).Run(data.Node.Rpc, data.Node.Id)
	}
}

// 服务注销通知
func (s *Server) ServerLogout(socket *ws.WebSocket, msg []byte) {
	var data RegisterMsg
	_ = json.Unmarshal(msg, &data)
	switch data.Node.Mode {
	case "match":
		utils.Println("匹配服务器离线")
		match.MatchService = nil
	case "gateway":
		utils.Println("网关离线")
		gateway.GatewayGrpc.Delete(data.Node.Id)
	}
}
