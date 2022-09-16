package routes

import (
	"github.com/gin-gonic/gin"
	"study-server/app/http/middleware"
	"study-server/bootstrap/config"
	"study-server/bootstrap/request"
	"study-server/routes"
)

type Route struct {
	Router     *gin.Engine                    // gin 路由
	controller func(*gin.Context) interface{} // 路由绑定的控制器
}

// 路由配置
func (r *Route) Handle() {

	// 全局中间件
	r.Router.Use(request.Handle)

	// 404
	r.Router.NoRoute(func(c *gin.Context) {
		if config.App.Server.Template {
			c.HTML(404, "404.tpl", gin.H{})
		} else {
			c.JSON(404, gin.H{"code": 404, "msg": "url not found"})
		}
	})

	switch config.App.Server.Mode {
	case "webserver":
		// web路由配置
		web := r.Router.Group("/")                       // 配置路由组
		web.Use((&middleware.Web{}).Handle)              // 使用中间件
		(&routes.Web{Route: web, Bind: r.bind}).Handle() // 关联子路由

		// 接口路由配置
		api := r.Router.Group("/api")
		api.Use((&middleware.Api{}).Handle)
		(&routes.Api{Route: api, Bind: r.bind}).Handle()

		// 后台管理接口路由配置
		admin := r.Router.Group("/admin")
		admin.Use((&middleware.Admin{}).Handle)
		admin.Use((&middleware.AdminRoleAuth{}).Handle)
		(&routes.Admin{Route: admin, Bind: r.bind}).Handle()
	default:
		// websocket路由
		ws := r.Router.Group("/ws")
		(&routes.WebSocket{Route: ws}).Handle()
	}

	// ... 声明其他路由文件
}
