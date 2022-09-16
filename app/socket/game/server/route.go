package server

import (
	"study-server/bootstrap/udp/server"
)

func (s *Server) Route() []*server.Route {
	return []*server.Route{
		{
			Action: "Move",
			Fun:    s.Move,
		},
	}
}