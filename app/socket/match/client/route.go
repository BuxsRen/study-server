package client

import ws "study-server/bootstrap/websocket/client"

// 事件路由
func (s *Server) Route() []*ws.Route {
	return []*ws.Route{
		{
			Action: "ServerRegister",
			Fun:    s.ServerRegister,
		}, {
			Action: "ServerLogout",
			Fun:    s.ServerLogout,
		},
	}
}