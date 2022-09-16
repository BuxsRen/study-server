package grpc

import (
	"context"
	"encoding/json"
	"study-server/app/libs/utils"
	"study-server/app/socket/gateway/gateway"
	pg "study-server/bootstrap/grpc/proto/gateway"
)

type Msg struct {
	Action  string `json:"action"`
	ToId    string `json:"toId"`
	FromId  string `json:"fromId"`
	Content string `json:"content"`
}

var _ pg.GatewayGrpcServer = (*Server)(nil)

// 其他网关用户登录消息事件
func (s *Server) OnLogin(_ context.Context, msg *pg.Msg) (*pg.Res, error) {
	if msg.Gateway != gateway.DeviceNumber { // 排除当前网关发出的
		c, ok := gateway.WebSocket.GetClient(msg.FromId)
		if ok { // 在本网关有登录该客户端，使其强制下线
			gateway.WebSocket.ForceLogout(c)
		}
	}
	utils.Println("[gateway] -> 用户登录", msg.FromId, msg.Gateway)
	return &pg.Res{}, nil
}

// 其他网关用户退出消息事件
func (s *Server) OnLogout(_ context.Context, msg *pg.Msg) (*pg.Res, error) {
	utils.Println("[gateway] -> 用户退出", msg.FromId, msg.Gateway)
	return &pg.Res{}, nil
}

// 其他网关用户加入区域消息事件
func (s *Server) OnAreaJoin(_ context.Context, msg *pg.Msg) (*pg.Res, error) {
	b, _ := json.Marshal(&Msg{Action: "AreaJoin", ToId: msg.ToId, FromId: msg.FromId, Content: msg.Content})
	_ = gateway.WebSocket.Send(msg.ToId, b)
	utils.Println("[gateway] -> 用户加入区域", msg.FromId, msg.Gateway, msg.Content)
	return &pg.Res{}, nil
}

// 其他网关退出区域消息事件
func (s *Server) OnAreaExit(_ context.Context, msg *pg.Msg) (*pg.Res, error) {
	b, _ := json.Marshal(&Msg{Action: "AreaExit", ToId: msg.ToId, FromId: msg.FromId, Content: msg.Content})
	_ = gateway.WebSocket.Send(msg.ToId, b)
	utils.Println("[gateway] -> 用户退出区域", msg.FromId, msg.Gateway, msg.Content)
	return &pg.Res{}, nil
}

// 其他网关用户发送消息事件
func (s *Server) OnSendMsg(_ context.Context, msg *pg.Msg) (*pg.Res, error) {
	b, _ := json.Marshal(&Msg{Action: "SendMsg", ToId: msg.ToId, FromId: msg.FromId, Content: msg.Content})
	_ = gateway.WebSocket.Send(msg.ToId, b)
	utils.Println("[gateway] -> 用户消息", msg.FromId, "=>", msg.ToId, msg.Content)
	return &pg.Res{}, nil
}
