// 匹配服务器GRPC服务端
package grpc

import (
	"context"
	"errors"
	"study-server/app/socket/match/game"
	"study-server/app/socket/match/match"
	"study-server/bootstrap/area"
	pm "study-server/bootstrap/grpc/proto/match"
	"sync"
)

var _ pm.MatchGrpcServer = (*Server)(nil)

// 改变区域/随机加入一个区域
func (s *Server) ChangeArea(ctx context.Context, req *pm.Req) (*pm.Res, error) {
	chanRes := make(chan area.Res)
	match.UserArea.ChangeArea(&area.Client{
		Id:      req.Id,
		Name:    req.Gateway,
		ChanRes: &chanRes,
	})
	// 等待回执消息
	select {
	case res, ok := <-chanRes:
		close(chanRes)
		if !ok {
			return &pm.Res{}, nil
		}
		gs, e := game.GetGameService(res.Area) // 如果未连接到游戏服务器则拿不到数据
		var url, id string
		if e == nil {
			url = gs.Url
			id = gs.Id
		}
		return &pm.Res{
			Action:  res.Action,
			Id:      res.Id,
			Area:    res.Area,
			Group:   res.Group,
			GameUrl: url,
			GameId:  id,
		}, res.Err
	}
}

// 加入到指定区域
func (s *Server) JoinArea(ctx context.Context, req *pm.Req) (*pm.Res, error) {
	chanRes := make(chan area.Res)
	match.UserArea.JoinArea(&area.Client{
		Id:      req.Id,
		Name:    req.Gateway,
		Area:    req.Area,
		ChanRes: &chanRes,
	})
	// 等待回执消息
	select {
	case res, ok := <-chanRes:
		close(chanRes)
		if !ok {
			return &pm.Res{}, nil
		}
		gs, e := game.GetGameService(res.Area)
		var url, id string
		if e == nil {
			url = gs.Url
			id = gs.Id
		}
		return &pm.Res{
			Action:  res.Action,
			Id:      res.Id,
			Area:    res.Area,
			Group:   res.Group,
			GameUrl: url,
			GameId:  id,
		}, res.Err
	}
}

// 退出区域
func (s *Server) ExitArea(ctx context.Context, req *pm.Req) (*pm.Res, error) {
	chanRes := make(chan area.Res)
	match.UserArea.ExitArea(&area.Client{
		Id:      req.Id,
		ChanRes: &chanRes,
	})
	// 等待回执消息
	select {
	case res, ok := <-chanRes:
		close(chanRes)
		if !ok {
			return &pm.Res{}, nil
		}
		gs, e := game.GetGameService(res.Area)
		var url, id string
		if e == nil {
			url = gs.Url
			id = gs.Id
		}
		return &pm.Res{
			Action:  res.Action,
			Id:      res.Id,
			Area:    res.Area,
			Group:   res.Group,
			GameUrl: url,
			GameId:  id,
		}, res.Err
	}
}

// 获取用户区域内信息(获取区域内所有用户id)
func (s *Server) GetAreaInfo(ctx context.Context, req *pm.Req) (*pm.List, error) {
	var list *sync.Map
	var e error
	var group string
	var areaName string
	var gameUrl string
	var gameId string
	if req.Id == "" && req.Area == "" {
		return &pm.List{}, errors.New("用户ID和区域ID不能都为空")
	}
	if req.Id != "" {
		info, e := match.UserArea.GetInfoById(req.Id) // 获取用户所在的区域
		if e != nil {
			return &pm.List{}, e
		}
		list, e = match.UserArea.GetAreaInfo(info.Group, info.Area) // 获取用户所在区域所有的客户端
		if e != nil {
			return &pm.List{}, e
		}
		group = info.Group
		areaName = info.Area
		gs, e := game.GetGameService(info.Area)
		if e == nil {
			gameUrl = gs.Url
			gameId = gs.Id
		}
	}
	if req.Area != "" {
		list, e = match.UserArea.GetAreaInfo("area", req.Area)
		if e != nil {
			return &pm.List{}, e
		}
		areaName = req.Area
		group = "area"
	}
	var ids []*pm.Node
	list.Range(func(key, value interface{}) bool {
		v := value.(*area.Client)
		ids = append(ids, &pm.Node{
			Id:      v.Id,
			Gateway: v.Name,
		})
		return true
	})
	return &pm.List{Id: ids, Area: areaName, Group: group, GameId: gameId, GameUrl: gameUrl}, e
}

// 获取用户信息
func (s *Server) GetUserInfo(ctx context.Context, req *pm.Req) (*pm.Res, error) {
	info, e := match.UserArea.GetInfoById(req.Id)
	if e != nil {
		return nil, e
	}
	gs, e := game.GetGameService(info.Area)
	var url, id string
	if e == nil {
		url = gs.Url
		id = gs.Id
	}
	return &pm.Res{Id: info.Id,
		Group:   info.Group,
		Area:    info.Area,
		Gateway: info.Name,
		GameUrl: url,
		GameId:  id,
	}, nil
}