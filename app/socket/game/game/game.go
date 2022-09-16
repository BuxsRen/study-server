package game

import (
	"fmt"
	"study-server/app/libs/utils"
	"study-server/bootstrap/area"
	"study-server/bootstrap/udp/server"
)

// 游戏服务器唯一编号
var DeviceNumber string
var GameArea *area.Area
var GameServer *server.UdpServer

func init() {
	DeviceNumber = "game_" + utils.GetRandString(5) + fmt.Sprintf("%v", utils.Rand(1000, 9999))
	GameArea = area.New(10, DeviceNumber, false) // 创建一个区域管理器，每个区域最大可容纳10人，用于网关客户端登录加入该区域
}
