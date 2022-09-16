package models

import (
	"study-server/app/libs/utils"
)

// 聊天记录表
type Message struct{}

func (m *Message) TableName() string {
	return "message"
}

// 获取未读消息记录数
func (m *Message) GetUnReadMessageCount(uid, from interface{}) int64 {
	list, e := DB(m).Fields("id", "from_id", "users_id", "content", "create_at").Where("users_id", uid).Where("from_id", from).Where("is_read", 0).Count()
	if e != nil {
		return 0
	}
	return list
}

// 追加聊天记录
func (m *Message) Push(uid, from, to, content interface{}, read interface{}) int64 {
	id, e := DB(m).Data(map[string]interface{}{
		"users_id":  uid,
		"from_id":   from,
		"to_id":     to,
		"content":   content,
		"is_read":   read,
		"create_at": utils.GetTime(),
	}).InsertGetId()
	if e != nil {
		return 0
	}
	return id
}

// 获取聊天记录
func (m *Message) GetMessageList(data map[string]interface{}) *PageData {
	table := DB(m)
	id := data["id"]
	table = table.Where("users_id", data["_id"]).Where("create_at", ">=", utils.GetTime()-86400*7).Where(func() {
		table.Where("from_id", id).OrWhere("to_id", id)
	}).Fields("id", "is_read", "from_id", "users_id", "content", "create_at")
	delete(data, "id")
	model := Model{Model: m, Table: table, Data: data, Order: "id", OrderType: "desc"}
	return model.PageSearch()
}

// 消息至已读
func (m *Message) MessageToRead(data map[string]interface{}) {
	_, _ = DB(m).Where("users_id", data["_id"]).Where("from_id", data["id"]).Data(map[string]interface{}{"is_read": 1}).Update()
}