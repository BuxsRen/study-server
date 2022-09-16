package api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"study-server/app/libs/code"
	"study-server/app/libs/email"
	"study-server/app/libs/encry"
	"study-server/app/libs/redis"
	"study-server/app/libs/upload"
	"study-server/app/libs/utils"
	"study-server/app/libs/verify"
	"study-server/app/models"
	"study-server/bootstrap/config"
	"study-server/bootstrap/request"
)

// UserController 用户相关接口控制器
type UserController struct{}

// Login 用户登录
func (uc *UserController) Login(c *gin.Context) interface{} {
	data := c.GetStringMap("_data")

	verify.VerifyData(data, map[string]string{
		"username": "required|string|email",
		"code":     "required|string|between:1,5",
		"password": "required|string|alpha_dash|between:5,19",
		"repeat":   "required|string|alpha_dash|between:5,19",
	})

	u := models.User{}
	user := u.CheckUser(data)

	if user == nil {
		return utils.NO("用户名或密码错误", -1, nil)
	}

	if user["is_ban"].(int64) == 1 {
		return utils.NO("该用户禁止登录", -38, nil)
	}

	delete(user, "delete_at")
	delete(user, "update_at")
	delete(user, "is_ban")

	user["token"] = encry.EncryptToken(user["id"].(string), 86400, user)
	return utils.OK("登录成功", user)
}

// Register 注册
func (uc *UserController) Register(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	type Param struct {
		Username string `json:"username" bind:"required,string,email"`
		Password string `json:"password" bind:"required,string,between=5:19"`
		Repeat   string `json:"repeat" bind:"required,string,between=5:19""`
		Code     string `json:"code" bind:"required,string,between=1:5"`
	}

	param, e := verify.Bind(req.Param, Param{})
	if e != nil {
		return utils.NO(e.Error(), -1, nil)
	}

	if param.Password != param.Repeat {
		return utils.NO("两次密码不一致", -1, nil)
	}

	rdb := redis.Redis{}
	if strings.ToUpper(rdb.Get("Emali_Code_"+param.Username)) != strings.ToUpper(param.Code) {
		return utils.NO("验证码不正确", -1, nil)
	}

	u := models.User{}
	user := u.CheckUserName(param.Username)
	if user != nil {
		return utils.NO("该账号已被注册", -1, nil)
	}

	u.Create(req.Param)
	return utils.OK("注册成功", nil)
}

// Emali 发送邮箱验证码
func (uc *UserController) Emali(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	verify.VerifyData(req.Param, map[string]string{
		"username": "required|string|email",
		"code":     "required|string|between:1,5",
		"id":       "required|string",
	})

	rdb := redis.Redis{}
	verify := rdb.Get("Verify_Code_" + req.Param["id"].(string))

	if strings.ToUpper(verify) != strings.ToUpper(req.Param["code"].(string)) {
		return utils.NO("验证码不正确", -1, nil)
	}

	em := email.New()
	em.SetTitle("Hello，注册")
	em.SetToEmail([]string{req.Param["username"].(string)})

	s, e := (&code.Code{}).CreateStrVerifyCode()
	if e != nil {
		return utils.NO("邮箱验证码生成失败", -1, nil)
	}

	em.SetBody("<h1>您的验证码是:" + strings.ToUpper(s.Str) + "<br>5分钟内有效</h1>")
	e = rdb.Setex("Email_Code_"+req.Param["username"].(string), s.Str, 500)
	if e != nil {
		return utils.NO("redis error", -1, nil)
	}

	e = em.SendMail()
	if e != nil {
		return utils.NO("邮箱验证码发送失败", -1, nil)
	}

	return utils.OK("邮箱验证码发送成功", nil)
}

// List 用户列表
func (uc *UserController) List(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	u := models.User{}
	list := u.List(req.Param)
	return utils.OK("获取成功", list)
}

// Info 获取用户信息
func (uc *UserController) Info(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	verify.VerifyData(req.Param, map[string]string{
		"id": "required|numeric",
	})

	u := models.User{}
	info := u.Info(req.Param)

	if info["is_ban"].(int64) == 1 {
		return utils.NO("该用户禁止登录", -38, nil)
	}
	return utils.OK("获取成功", info)
}

// GetMessageList 获取聊天记录
func (uc *UserController) GetMessageList(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	verify.VerifyData(req.Param, map[string]string{
		"id": "required|numeric",
	})

	m := models.Message{}
	list := m.GetMessageList(req.Param)
	return utils.OK("获取成功", list)
}

// UpdateInfo 更新资料
//func (uc *UserController) UpdateInfo(c *gin.Context) interface{} {
//	id := request.GetInput(c, "id")
//	req := request.GetAllData(c)
//	u := models.User{}
//	return utils.OK("资料更新成功", u.Update(id, req.Param))
//}

// Upload 上传
func (uc *UserController) Upload(c *gin.Context) interface{} {
	req := request.GetAllData(c)

	verify.VerifyData(req.Param, map[string]string{
		"format": "required|string",
	})

	file, err := c.FormFile("file")
	if err != nil {
		utils.ExitError("请选择文件上传", -1)
	}

	f := upload.Upload{File: file, Format: req.Param["format"].(string)}
	path := f.Upload()
	prefix := config.App.Other.PublicPrefix
	if prefix == "" {
		utils.ExitError("未设置 public_prefix", -1)
	}

	url := config.App.Server.Url
	return utils.OK("上传成功", url+prefix+path)
}
