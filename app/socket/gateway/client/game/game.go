package game

import (
	"context"
	"study-server/app/socket/gateway/gateway"
	pg "study-server/bootstrap/grpc/proto/game"
)

// 设置游戏服务玩家的区域
func (g *Game) SetArea(id, group, area, gameId string) (*pg.Res, error) {
	s, e := g.GetGame(gameId)
	if e != nil {
		return &pg.Res{}, e
	}
	_, _ = (*s.server).SetArea(context.Background(), &pg.Req{Id: id, Group: group, Area: area, Gateway: gateway.DeviceNumber})
	return &pg.Res{}, nil
}

// 删除游戏服务玩家的区域
func (g *Game) DelArea(id string) (*pg.Res, error) {
	list := g.GetAllService()
	for _, v := range list {
		_, _ = (*v.server).DelArea(context.Background(), &pg.Req{Id: id})
	}
	return &pg.Res{}, nil
}