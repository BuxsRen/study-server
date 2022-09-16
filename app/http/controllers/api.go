package controllers

import (
	"github.com/gin-gonic/gin"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"study-server/app/libs/verify"
)

type ApiController struct{}

// 获取用户Token
func (api *ApiController) GetUserToken(c *gin.Context) interface{} {
	type Param struct {
		Id string `json:"id" bind:"required"`
	}
	param, _ := verify.Bind(c.GetStringMap("_data"), Param{})
	return utils.OK("获取成功", encry.EncryptToken(utils.ParamToString(param.Id), -1, nil))
}
