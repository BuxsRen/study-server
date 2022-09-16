package routes

import (
	"github.com/gin-gonic/gin"
	gateway "study-server/app/socket/gateway/server"
	manage "study-server/app/socket/manage/server"
	"study-server/app/socket/middleware"
	"study-server/bootstrap/config"
)

// WebSocket websocket路由
type WebSocket struct {
	Route *gin.RouterGroup // gin 路由
	// WebSocket 不需要绑定返回数据处理
}

// Handle 入口方法
func (ws *WebSocket) Handle() {
	switch config.App.Server.Mode {
	case "gateway":
		ws.Route.Use((&middleware.Gateway{}).Handle).GET("gateway", gateway.Handle) // 网关服务器
	case "manage":
		ws.Route.Use((&middleware.Manage{}).Handle).GET("manage", manage.Handle) // 管理服务器
	}
}
