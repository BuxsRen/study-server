package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
	"time"
)

// 逐一将路由绑定到控制器 (gin路由(GET,POST...)，路由地址，指向控制器)
func (r *Route) bind(method func(string, ...gin.HandlerFunc) gin.IRoutes, url string, controller func(*gin.Context) interface{}) {
	if config.App.Server.Debug {
		//m := utils.GetSubstr(runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(), "github.com/gin-gonic/gin.(*RouterGroup).", "-fm")
		c := utils.GetSubstr(runtime.FuncForPC(reflect.ValueOf(controller).Pointer()).Name(), "/", "-fm")
		defer fmt.Printf("\u001B[%dm[controllers] --> %s \u001B[0m\n\n", 34, c)
	}
	c := &Route{controller: controller}
	method(url, c.handler) // 绑定路由地址和控制器到gin路由
}

// 执行控制器，获取结果并响应
func (r *Route) handler(c *gin.Context) {
	s := utils.GetNow()
	data := r.controller(c)
	c.Header("Rep-Time", time.Since(s).String()) // 控制器处理消耗时间
	t := fmt.Sprintf("%T", data)
	switch t {
	case "string": // 字符串处理 return "test"
		c.String(200, data.(string))
	case "helper.Views": // 模板处理 return helper.View("404.tpl",nil)
		temp := data.(helper.Views)
		c.HTML(200, temp.Template, temp.Data)
	case "<nil>":
		c.String(200, "")
	default: // 默认json处理 return map[string]interface{}{"code":200}
		c.JSON(200, data)
	}
}
