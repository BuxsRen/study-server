package gateway

import (
	"context"
	pg "study-server/bootstrap/grpc/proto/gateway"
)

// 给网关推送其他网关客户端登录消息
func (m *Gateway) OnLogin(id, gateway string) (*pg.Res, error) {
	return (*m.service).OnLogin(context.Background(), &pg.Msg{
		FromId:  id,
		Gateway: gateway,
	})
}

// 给网关推送其他网关客户端离线消息
func (m *Gateway) OnLogout(id, gateway string) (*pg.Res, error) {
	return (*m.service).OnLogout(context.Background(), &pg.Msg{
		FromId:  id,
		Gateway: gateway,
	})
}

// 给网关推送其他网关客户端加入区域消息
func (m *Gateway) OnAreaJoin(id, toId string) (*pg.Res, error) {
	return (*m.service).OnAreaJoin(context.Background(), &pg.Msg{
		FromId: id,
		ToId:   toId,
	})
}

// 给网关推送其他网关客户端退出区域消息
func (m *Gateway) OnAreaExit(id, toId string) (*pg.Res, error) {
	return (*m.service).OnAreaExit(context.Background(), &pg.Msg{
		FromId: id,
		ToId:   toId,
	})
}

// 给网关推送其他网关客户端消息
func (m *Gateway) OnSendMsg(id, toId, content string) (*pg.Res, error) {
	return (*m.service).OnSendMsg(context.Background(), &pg.Msg{
		FromId:  id,
		ToId:    toId,
		Content: content,
	})
}
