package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study-server/app/socket/manage/manage"
	ws "study-server/bootstrap/websocket/server"
)

// 管理服务器
// 负责节点注册,节点获取节点,节点上线推送等
type ManageServer struct{}

var _ ws.WebSocketInterface = (*ManageServer)(nil)

// WebSocket客户端接入处理方法
func Handle(c *gin.Context) {
	client, e := manage.WebSocket.Bind(c, new(ManageServer))
	if e != nil {
		fmt.Println("[manage]", e)
		return
	}
	_ = manage.WebSocket.Start(client)
}