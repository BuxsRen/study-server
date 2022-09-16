package server

import (
	ws "study-server/bootstrap/websocket/server"
)

// WebSocket 事件路由
func (ms *ManageServer) Route() []*ws.Route {
	return []*ws.Route{}
}
