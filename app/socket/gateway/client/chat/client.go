package chat

import (
	"errors"
	"study-server/app/libs/utils"
	gc "study-server/bootstrap/grpc/client"
	pc "study-server/bootstrap/grpc/proto/chat"
	"sync"
)

type Chat struct {
	service *pc.ChatGrpcClient
}

var ChatGrpc sync.Map // *Chat

// 网关服务器连接聊天grpc服务器
func (m *Chat) Run(grpcUrl, id string) {
	utils.Println("[gateway][<=>][chat] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	cgc := pc.NewChatGrpcClient(conn)
	c := &Chat{service: &cgc}
	ChatGrpc.Store(id, c)
}

// 随机取一个在线的聊天服务器
func GetChatService() (*Chat, error) {
	var c *Chat
	ChatGrpc.Range(func(key, value interface{}) bool {
		c = value.(*Chat)
		return true
	})
	if c != nil {
		return c, nil
	}
	return nil, errors.New("网关 -> 未连接到 -> 聊天服务器")
}