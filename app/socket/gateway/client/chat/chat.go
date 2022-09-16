package chat

import (
	"context"
	"study-server/app/socket/gateway/gateway"
	pc "study-server/bootstrap/grpc/proto/chat"
)

// 给其他网关推送该网关客户端上线通知
func (c *Chat) Login(id string) (*pc.Res, error) {
	return (*c.service).Login(context.Background(), &pc.Req{Id: id, Gateway: gateway.DeviceNumber})
}

// 给其他网关推送该网关客户端离线通知
func (c *Chat) Logout(id string) (*pc.Res, error) {
	return (*c.service).Logout(context.Background(), &pc.Req{Id: id, Gateway: gateway.DeviceNumber})
}

// 给区域内所有用户推送加入通知
func (c *Chat) JoinArea(id, area string) (*pc.Res, error) {
	return (*c.service).JoinArea(context.Background(), &pc.Req{Id: id, Gateway: gateway.DeviceNumber, OldArea: area})
}

// 给区域内所有用户推送退出消息
func (c *Chat) ExitArea(id, area string) (*pc.Res, error) {
	return (*c.service).ExitArea(context.Background(), &pc.Req{Id: id, Gateway: gateway.DeviceNumber, Area: area})
}

// 给指定用户推送消息
func (c *Chat) SendMsg(id, toId, content string) (*pc.Res, error) {
	return (*c.service).SendMsg(context.Background(), &pc.ReqMsg{
		FromId:  id,
		Content: content,
		ToId:    toId,
	})
}

// 给指定区域广播消息
func (c *Chat) SendMsgArea(id, content string) (*pc.Res, error) {
	return (*c.service).SendMsgArea(context.Background(), &pc.ReqMsg{
		FromId:  id,
		Content: content,
	})
}
