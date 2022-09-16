package models

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
)

// User 用户表
type User struct{}

func (u *User) TableName() string {
	return "users"
}

// CheckUser 根据用户名和密码查询用户
func (u *User) CheckUser(data map[string]interface{}) gorose.Data {
	res, e := DB(u).Where("username", data["username"]).Where("password", encry.Password(data["password"].(string))).Where("delete_at", 0).First()
	if e != nil {
		return nil
	}
	if data["product"] != nil && data["manufacturer"] != nil {
		device := utils.ParamToString(data["manufacturer"]) + " " + utils.ParamToString(data["product"])
		_, _ = DB(u).Where("id", res["id"]).Data(map[string]interface{}{"device": device, "last_login": utils.GetTime()}).Update()
	}
	return res
}

// CheckUserName 根据用户名
func (u *User) CheckUserName(username string) gorose.Data {
	res, e := DB(u).Where("username", username).Where("delete_at", 0).First()
	if e != nil {
		return gorose.Data{}
	}
	return res
}

// Create 注册用户
func (u *User) Create(data map[string]interface{}) int64 {
	ins := map[string]interface{}{
		"username":  data["username"],
		"nickname":  "Hello" + fmt.Sprintf("%v", utils.Rand(100000, 999999)),
		"password":  encry.Password(data["password"].(string)),
		"birthday":  utils.GetTime(),
		"create_at": utils.GetTime(),
	}
	id, e := DB(u).Data(ins).Insert()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	return id
}

// List 用户列表
func (u *User) List(data map[string]interface{}) *PageData {
	table := DB(u).Fields("id", "username", "nickname", "avatar", "is_ban", "cover", "star", "sex", "birthday", "identity", "address")
	m := &Model{Model: u, Table: table, Data: data}
	res := m.PageSearch()
	msg := Message{}
	for k, v := range res.Items {
		res.Items[k]["un_read_count"] = msg.GetUnReadMessageCount(data["_id"], v["id"])
	}
	return res
}

// Info 用户信息
func (u *User) Info(data map[string]interface{}) gorose.Data {
	info, e := DB(u).Where("id", data["id"]).Fields("id", "username", "nickname", "create_at", "avatar", "cover", "star", "sex", "birthday", "identity", "is_ban", "info", "device", "last_login", "address").First()
	if data["id"] == utils.ParamToString(data["_id"]) && data["product"] != nil && data["manufacturer"] != nil {
		device := utils.ParamToString(data["manufacturer"]) + " " + utils.ParamToString(data["product"])
		_, _ = DB(u).Where("id", data["id"]).Data(map[string]interface{}{"device": device, "last_login": utils.GetTime()}).Update()
	}
	if e != nil {
		return gorose.Data{}
	}
	return info
}
