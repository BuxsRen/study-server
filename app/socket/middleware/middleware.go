package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
)

type Middleware struct {
	C     *gin.Context
	Token string
}

func New(c *gin.Context) (*Middleware, error) {
	m := &Middleware{C: c}
	m.Token = m.C.Query("token")
	if m.Token == "" {
		m.C.JSON(401, gin.H{"code": 401, "msg": "缺少必要的参数，无法验证身份"})
		m.C.Abort()
		return m, errors.New("缺少必要的参数，无法验证身份")
	}
	return m, nil
}

// 网关登录 -> 验证用户token
func (m *Middleware) CheckToken() bool {
	data := encry.DecryptToken(m.Token)
	if data == nil {
		m.C.JSON(-99, gin.H{"code": -99, "msg": "login timeout"})
		return false
	} else {
		m.C.Set("_id", utils.ParamToString(data["id"]))
		m.C.Set("_name", utils.ParamToString(data["username"]))
		m.C.Set("_group", "gateway")
		return true
	}
}

// 管理端登录 -> 验证授权
func (m *Middleware) CheckAdmin() bool {
	info := encry.DecryptToken(m.Token)
	if info == nil {
		m.C.JSON(-500, gin.H{"code": 500, "msg": "token解析失败"})
		return false
	}
	m.C.Set("_id", utils.ParamToString(info["id"]))
	m.C.Set("_group", utils.ParamToString(info["group"]))
	m.C.Set("_rpc", utils.ParamToString(info["rpc"]))
	m.C.Set("_udp", utils.ParamToString(info["udp"]))
	return true
}
