package match

import (
	"study-server/app/libs/utils"
	"fmt"
	"study-server/bootstrap/area"
)

// 聊天服务器唯一编号
var DeviceNumber string

var UserArea *area.Area

func init() {
	DeviceNumber = "match_" + utils.GetRandString(5) + fmt.Sprintf("%v", utils.Rand(1000, 9999))
	UserArea = area.New(10, DeviceNumber, false) // 创建一个区域管理器，每个区域最大可容纳10人，用于网关客户端登录加入该区域
}