package grpc

import (
	"context"
	"study-server/app/libs/utils"
	"study-server/app/socket/game/game"
	"study-server/bootstrap/area"
	pg "study-server/bootstrap/grpc/proto/game"
)

var _ pg.GameGrpcServer = (*Server)(nil)

func (s *Server) SetArea(ctx context.Context, req *pg.Req) (*pg.Res, error) {
	utils.Println("用户加入游戏区域", req)
	chanRes := make(chan area.Res, 1)
	game.GameArea.CustomArea(&area.Client{Id: req.Id, Group: req.Group, Area: req.Area, Name: req.Gateway, ChanRes: &chanRes})
	select {
	case msg, ok := <-chanRes:
		if !ok {
			return &pg.Res{}, nil
		}
		return &pg.Res{}, msg.Err
	}
}

func (s *Server) DelArea(ctx context.Context, req *pg.Req) (*pg.Res, error) {
	utils.Println("用户退出游戏区域", req)
	chanRes := make(chan area.Res, 1)
	game.GameArea.ExitArea(&area.Client{Id: req.Id, ChanRes: &chanRes})
	select {
	case msg, ok := <-chanRes:
		if !ok {
			return &pg.Res{}, nil
		}
		return &pg.Res{}, msg.Err
	}
}