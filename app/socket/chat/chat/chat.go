package chat

import (
	"fmt"
	"study-server/app/libs/utils"
)

// 聊天服务器唯一编号
var DeviceNumber string

func init() {
	DeviceNumber = "chat_" + utils.GetRandString(5) + fmt.Sprintf("%v", utils.Rand(1000, 9999))
}
