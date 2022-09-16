package middleware

import "github.com/gin-gonic/gin"

type Manage struct {}

func (manage *Manage) Handle(c *gin.Context)  {
	m,e := New(c)
	if e != nil {
		return
	}
	if m.CheckAdmin() {
		m.C.Next()
	} else {
		m.C.Abort()
	}
}