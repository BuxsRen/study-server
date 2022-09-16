package manage

import (
	"fmt"
	"study-server/app/libs/utils"
	ws "study-server/bootstrap/websocket/server"
)

// 网关唯一编号
var DeviceNumber string

// WebSocket服务
var WebSocket *ws.WebSocket

func init() {
	DeviceNumber = "manage_" + utils.GetRandString(5) + fmt.Sprintf("%v", utils.Rand(1000, 9999)) // 生成一个服务唯一id
	WebSocket = ws.New(DeviceNumber, 100)                                                         // 创建一个Websocket服务端
}