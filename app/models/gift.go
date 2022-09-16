package models

import (
	"github.com/gohouse/gorose/v2"
	"study-server/app/libs/utils"
)

// 抽奖活动表
type Gift struct{}

func (g *Gift) TableName() string {
	return "gifts"
}

// 获取抽奖活动列表
func (g *Gift) GetList(aid interface{}, field ...string) []gorose.Data {
	list, e := DB(g).Where("activitys_id", aid).Where("delete_at", 0).Fields(field...).Get()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}
	return list
}
