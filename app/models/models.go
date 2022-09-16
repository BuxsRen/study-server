// 数据库操作基础类
package models

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"strconv"
	"strings"
	"study-server/app/libs/utils"
	"study-server/bootstrap/database"
)

// DB 获取表操作实例
/**
 *@Example:
	type User struct{}
	func (u *User) TableName() string {
		return "users"
	}
	func (u *User) GetUserInfo(id int) gorose.Data {
		res,_ := DB(u).Where("id",id).First()
		return res
	}
*/
func DB(table interface{}) gorose.IOrm {
	return database.GetPoolClinet().NewOrm().Table(table)
}

func DbSession(table interface{}) gorose.ISession {
	return database.GetPoolClinet().NewSession().Bind(table)
}

var (
	page        = 1                                                                               // 分页默认开始页数
	limit       = 20                                                                              // 分页默认每页条数
	selectBlack = []string{"page", "limit", "file", "access_token", "token", "x-token", "images"} // 不参与筛选的参数 下划线开头的也不会参与
	selectVague = []string{"username", "name", "nickname", "remarks", "title", "school"}          // 参与模糊查询的参数
)

type Model struct {
	Table     gorose.IOrm
	Data      map[string]interface{} // 需要筛选的数据
	Order     string                 // 排序字段
	OrderType string                 // 排序方式 asc,desc
	Search    map[string]interface{} // 搜索关键词
	Model     interface{}            // 操作的model
}

type PageData struct {
	Count int64         `json:"count"`
	Items []gorose.Data `json:"items"`
}

// PageSearch 数据筛选/分页
/**
 *@Example:
	m := Model{
		Model:u,
		Table: DB(u),
		Data: data,
		Search: map[string]interface{}{
			"title":"测试",
		},
	}
	return m.PageSearch()
*/
func (m *Model) PageSearch() *PageData {
	count := DB(m.Model)
	search := DB(m.Model)

	for k, v := range m.Data { // 筛选
		if k[0:1] == "_" { // 忽略 _ 开头的参数
			continue
		}
		if !utils.InArray(k, selectBlack) && !utils.InArray(k, selectVague) { // 不在黑名单 且 不在 模糊搜索名单的键
			if fmt.Sprintf("%T", v) != "[]string" && fmt.Sprintf("%T", v) != "map[string]interface{}" { // 条件查询
				m.Table.Where(k, v)
				count.Where(k, v)
				search.Where(k, v)
			}
		} else if utils.InArray(k, selectVague) { // 模糊筛选
			item := utils.ParamToString(v)
			m.Table.Where(k, "like", "%"+item+"%")
			count.Where(k, "like", "%"+item+"%")
			search.Where(k, "like", "%"+item+"%")
		}
	}

	if fmt.Sprintf("%T", m.Search) != "map[]" { // 搜索 {"title":"关键字"};
		var result = false
		var key string
		for k, v := range m.Search {
			key = k
			_, res := search.Where(k, "like", "%"+utils.ParamToString(v)+"%").Get()
			if res == nil { // 搜索到结果直接退出，搜不到继续搜
				result = true
				m.Table.Where(k, "like", "%"+utils.ParamToString(v)+"%")
				count.Where(k, "like", "%"+utils.ParamToString(v)+"%")
				break
			}
		}
		if !result { // 搜索不到就把记录为空
			m.Table.Where(key, "@!->:)!<-@")
			count.Where(key, "@!->:)<-!@")
		}
	}

	if m.Data["_sort"] != nil { // 处理排序字段[接口专用] 格式 升序：_sort=+|id 或 降序：_sort=-|id
		sort := strings.Split(utils.ParamToString(m.Data["_sort"]), "|")
		if len(sort) == 2 {
			m.Order = sort[1]
			if sort[0] == "-" {
				m.OrderType = "desc"
			} else {
				m.OrderType = "asc"
			}
		}
	}

	num, _ := count.Count() // 统计

	if m.Data["page"] != nil {
		page, _ = strconv.Atoi(m.Data["page"].(string))
	} else {
		page = 1
	}
	if m.Data["limit"] != nil {
		limit, _ = strconv.Atoi(m.Data["limit"].(string))
	} else {
		limit = 20
	}
	m.Table.Offset((page - 1) * limit).Limit(limit) // 分页

	if m.Order != "" { // 排序
		if m.OrderType == "" {
			m.OrderType = "desc"
		}
		m.Table.Order(m.Order + " " + m.OrderType)
	}

	list, e := m.Table.Get()
	if e != nil {
		utils.ExitError(e.Error(), -1)
	}

	return &PageData{
		Count: num,
		Items: list,
	}
}
