package models

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"study-server/app/libs/encry"
	"study-server/app/libs/redis"
	"study-server/app/libs/utils"
)

// 活动表
type Activity struct{}

func (a *Activity) TableName() string {
	return "activitys"
}

// 获取活动列表
func (a *Activity) GetList(data map[string]interface{}) *PageData {
	table := DB(a).Fields("title", "start")
	m := Model{Model: a, Table: table, Data: data}
	return m.PageSearch()
}

// 获取活动详情
func (a *Activity) GetInfo(id interface{}) gorose.Data {
	rdb := redis.Redis{}
	key := fmt.Sprintf("Activitys_Info_Api_%s", id)
	rdb.CheckOut(key)
	res, e := DB(a).Where("id", id).First()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	if res["category"].(int64) == 3 {
		gift := Gift{}
		res["gift"] = gift.GetList(res["id"], "id", "name", "create_at")
	}
	str, _ := encry.JsonEncode(res)
	_ = rdb.Setex(key, str, 60)
	return res
}
