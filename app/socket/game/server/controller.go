package server

import (
	"study-server/app/socket/game/game"
	"study-server/bootstrap/udp/server"
)

func (s *Server) OnLogin(client *server.Client) {
	_ = client.Send([]byte(`{"action":"Login","message":"success","loginId":"` + client.Id + `"}`))
}

// 玩家移动消息处理
func (s *Server) Move(client *server.Client, msg []byte) {
	c, e := game.GameArea.GetInfoById(client.Id)
	if e != nil {
		_ = client.Send([]byte(`{"action":"Error","code":401,"msg":"无法感知用户所在区域"}`))
		return
	}
	_ = game.GameArea.SendAllInArea(msg, c.Group, c.Area)
}
