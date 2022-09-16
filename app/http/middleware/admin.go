package middleware

import (
	"github.com/gin-gonic/gin"
)

// 后台管理路由中间件
type Admin struct {
	c     *gin.Context
	id    int                    // 用户id
	token string                 // 用户token
	info  map[string]interface{} // 原始用户信息
}

// 入口方法
func (a *Admin) Handle(c *gin.Context) {
	admin := Admin{c: c}
	if !admin.checkUser() {
		admin.c.Abort()
	} else {
		admin.addParam()
		admin.c.Next()
	}
}

// 验证用户
func (a *Admin) checkUser() bool {
	return true
	//token := a.c.Request.Header.Get("X-Token")
	//if token == "" {
	//	token = a.c.Query("token")
	//}
	//rdb := redis.Redis{}
	//info := rdb.Get("Admin_" + token)
	//if info == "" {
	//	a.c.JSON(200, gin.H{
	//		"code": -99,
	//		"msg":  "还没登录或登录状态超时",
	//	})
	//	return false
	//}
	//data := make(map[string]interface{})
	//_ = json.Unmarshal([]byte(info), &data)
	//a.id = int(data["id"].(float64))
	//a.info = data
	//a.token = token
	//return true
}

// 追加参数
func (a *Admin) addParam() {
	//data := request.GetAllData(a.c)
	//data["_id"] = a.id
	//data["_token"] = a.token
	//data["_info"] = a.info
	//a.c.Set("_data", data)
}
