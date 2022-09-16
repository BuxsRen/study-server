package routes

import (
	"github.com/gin-gonic/gin"
	"study-server/app/http/controllers"
	"study-server/app/http/controllers/api"
	"study-server/bootstrap/config"
)

// Web web路由
type Web struct {
	Route *gin.RouterGroup                                                                           // gin 路由
	Bind  func(func(string, ...gin.HandlerFunc) gin.IRoutes, string, func(*gin.Context) interface{}) // 路由与控制器绑定：请求方式，路由，控制器
}

// Handle 入口方法
func (route *Web) Handle() {

	// 默认页
	route.Route.GET("/", func(c *gin.Context) {
		if config.App.Server.Template {
			c.HTML(200, "index.tpl", gin.H{})
		} else {
			c.String(200, "Hello World")
		}
	})

	route.Bind(route.Route.POST, "/api/get/user/token", (&controllers.ApiController{}).GetUserToken)

	route.Bind(route.Route.POST, "/api/register", (&api.UserController{}).Register)

	route.Bind(route.Route.POST, "/api/emali/code", (&api.UserController{}).Emali)

}
