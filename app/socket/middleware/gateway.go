package middleware

import (
	"github.com/gin-gonic/gin"
)

type Gateway struct {}

func (gw *Gateway) Handle(c *gin.Context)  {
	m,e := New(c)
	if e != nil {
		m.C.Abort()
		return
	}
	if m.CheckToken() {
		m.C.Next()
	} else {
		m.C.Abort()
	}
}

