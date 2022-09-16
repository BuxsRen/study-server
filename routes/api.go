package routes

import (
	"github.com/gin-gonic/gin"
	"study-server/app/http/controllers/api"
)

// Api 应用接口路由
type Api struct {
	Route *gin.RouterGroup                                                                           // gin 路由
	Bind  func(func(string, ...gin.HandlerFunc) gin.IRoutes, string, func(*gin.Context) interface{}) // 路由与控制器绑定：请求方式，路由，控制器
}

// Handle 入口方法
func (route *Api) Handle() {
	route.Bind(route.Route.GET, "/user/list", (&api.UserController{}).List)
	route.Bind(route.Route.GET, "/user/info", (&api.UserController{}).Info)
	route.Bind(route.Route.GET, "/message/list", (&api.UserController{}).GetMessageList)
	//route.Bind(route.Route.POST,"/user/update",(&api.UserController{}).UpdateInfo)
	route.Bind(route.Route.POST, "/upload", (&api.UserController{}).Upload)
}
