package client

import (
	"encoding/json"
	"study-server/app/libs/utils"
	"study-server/app/socket/match/game"
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
	case "game":
		utils.Println("游戏服务器离线")
		game.GameGrpc.Delete(data.Node.Id)
		gs, e := game.GetRandGameService()
		game.GameArea.Range(func(key, value interface{}) bool { // 从现有的区域中将已分配的掉线的游戏服务器进行剔除并赋予可用的服务器
			g := value.(*game.Game)
			if g.Id == data.Node.Id {
				if e == nil {
					game.GameArea.Store(key, gs)
				} else {
					game.GameArea.Delete(key)
				}
			}
			return true
		})
	}
}
