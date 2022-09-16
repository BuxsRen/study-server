package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study-server/app/socket/gateway/gateway"
	ws "study-server/bootstrap/websocket/server"
)

// 网关服务器
// 负责客户端接入，消息转发，推送，消息处理，客户端验证等
type GatewayServer struct{}

var _ ws.WebSocketInterface = (*GatewayServer)(nil)

// WebSocket客户端接入处理方法
func Handle(c *gin.Context) {
	client, e := gateway.WebSocket.Bind(c, new(GatewayServer))
	if e != nil {
		fmt.Println("[gateway]", e)
		return
	}
	_ = gateway.WebSocket.Start(client)
}
