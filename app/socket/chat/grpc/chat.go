package grpc

import (
	"context"
	"errors"
	"study-server/app/libs/utils"
	"study-server/app/socket/chat/client/gateway"
	"study-server/app/socket/chat/client/match"
	pc "study-server/bootstrap/grpc/proto/chat"
)

var _ pc.ChatGrpcServer = (*Server)(nil)

// 网关用户登录事件
func (s *Server) Login(ctx context.Context, req *pc.Req) (*pc.Res, error) {
	// 给所有网关推送该用户上线的消息，如果在用户在其他网关登录会被强制下线,忽略掉该用户已经登录的网关
	utils.Println("[chat] 用户登录 ->", req.Id, req.Area)
	list := gateway.GetAllGateway()
	for _, gw := range list { // 给已连接的所有网关推送用户登录消息
		_, _ = gw.OnLogin(req.Id, req.Gateway)
	}
	return &pc.Res{}, nil
}

// 网关用户退出事件
func (s *Server) Logout(ctx context.Context, req *pc.Req) (*pc.Res, error) {
	utils.Println("[chat] 用户退出 ->", req.Id, req.Area)
	list := gateway.GetAllGateway()
	for _, gw := range list { // 给已连接的所有网关推送用户离线消息
		_, _ = gw.OnLogout(req.Id, req.Gateway)
	}
	return &pc.Res{}, nil
}

// 加入区域，区域内消息广播
func (s *Server) JoinArea(ctx context.Context, req *pc.Req) (*pc.Res, error) {
	// 获取当前用户所在区域内所有用户Id
	utils.Println("[chat] 用户加入区域 ->", req.Id, req.Area)
	ms, e := match.GetMatchService()
	if e != nil {
		return &pc.Res{}, e
	}
	list, e := ms.GetAreaInfo(req.Id)
	if e != nil {
		return &pc.Res{}, e
	}
	// 给同区域所有用户推送该用户上线消息
	for _, v := range list.Id {
		gw, e := gateway.GetGateway(v.Gateway)
		if e != nil {
			continue
		}
		// 忽略自己
		if req.Id == v.Id {
			continue
		}
		_, _ = gw.OnAreaJoin(req.Id, v.Id)
	}
	return &pc.Res{}, nil
}

// 退出区域，区域内消息广播
func (s *Server) ExitArea(ctx context.Context, req *pc.Req) (*pc.Res, error) {
	// 获取当前用户所在区域内所有用户Id
	utils.Println("[chat] 用户退出区域 ->", req.Id, req.Area)
	ms, e := match.GetMatchService()
	if e != nil {
		utils.Println("[chat] 匹配服务器", e)
		return &pc.Res{}, e
	}
	list, e := ms.GetArea(req.Area)
	if e != nil {
		utils.Println("[chat] 获取区域信息", e)
		return &pc.Res{}, e
	}
	// 给同区域所有用户推送该用户上线消息
	for _, v := range list.Id {
		gw, e := gateway.GetGateway(v.Gateway)
		if e != nil {
			continue
		}
		// 忽略自己
		if req.Id == v.Id {
			continue
		}
		_, _ = gw.OnAreaExit(req.Id, v.Id)
	}
	return &pc.Res{}, nil
}

// 给指定用户发送消息
func (s *Server) SendMsg(ctx context.Context, req *pc.ReqMsg) (*pc.Res, error) {
	utils.Println("[chat] 用户消息 ->", req.FromId, "=>", req.ToId, " --->  ", req.Content)
	if req.FromId == req.ToId {
		return &pc.Res{}, errors.New("发送给自己的消息不出处理")
	}
	ms, e := match.GetMatchService()
	if e != nil {
		utils.Println(e)
		return &pc.Res{}, e
	}
	info, e := ms.GetUserInfo(req.ToId)
	if e != nil {
		return &pc.Res{}, e
	}
	gw, e := gateway.GetGateway(info.Gateway)
	if e != nil {
		return nil, e
	}
	_, e = gw.OnSendMsg(req.FromId, req.ToId, req.Content)
	return &pc.Res{}, e
}

// 用户所在区域内广播消息
func (s *Server) SendMsgArea(ctx context.Context, req *pc.ReqMsg) (*pc.Res, error) {
	ms, e := match.GetMatchService()
	if e != nil {
		return &pc.Res{}, e
	}
	list, e := ms.GetAreaInfo(req.FromId)
	if e != nil {
		utils.Println("[chat] 区域消息广播失败 ->", req.FromId, " | ", e, " | ", req.Content)
		return &pc.Res{}, e
	}
	for _, v := range list.Id {
		if req.FromId == v.Id {
			continue
		} // 排除自己
		gw, e := gateway.GetGateway(v.Gateway)
		if e != nil {
			return nil, e
		}
		_, _ = gw.OnSendMsg(req.FromId, v.Id, req.Content)
	}
	return &pc.Res{}, nil
}