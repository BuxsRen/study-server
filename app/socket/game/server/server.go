// 即时(战斗)服务器
package server

import (
	"study-server/app/socket/game/game"
	"study-server/bootstrap/udp/server"
)

type Server struct{}

// UDP Server端
func Star() {
	game.GameServer = server.New(new(Server))
	game.GameArea.UseSend(game.GameServer.Send)
	game.GameServer.Star(1024)
}
