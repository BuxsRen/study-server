package server

import (
	"encoding/json"
	"study-server/app/libs/utils"
	cc "study-server/app/socket/gateway/client/chat"
	"study-server/app/socket/gateway/client/game"
	"study-server/app/socket/gateway/client/match"
	ws "study-server/bootstrap/websocket/server"
)

// 客户端登录
func (gws *GatewayServer) Login(client *ws.Client, force bool) {
	utils.Println("[gateway] 用户登录", client.Id, force)
	// 正常上线，该网关没有发送强制挤掉原设备的情况
	if !force {
		cs, e := cc.GetChatService()
		if e != nil {
			return
		}
		// 给其他网关推送上线通知。如果其他网关有当前客户端在线，使其强制下线
		_, e = cs.Login(client.Id)
		if e != nil {
			utils.Println("[gateway] 用户登录转发至聊天服务器失败", e)
		}
	}
	gws.ChangeArea(client, nil)
}

// 客户端退出
func (gws *GatewayServer) Logout(client *ws.Client) {
	utils.Println("[gateway] 用户退出", client.Id)
	// 退出区域
	ms, e := match.GetMatchService()
	if e != nil {
		return
	}
	res, _ := ms.ExitArea(client.Id)
	g := &game.Game{}
	_, _ = g.DelArea(client.Id)
	if client.Force {
		utils.Println("[gateway] 用户被迫下线", client.Id)
	} else {
		utils.Println("[gateway] 用户下线", client.Id, res.GameId)
	}
	// 给其他网关推送下线通知
	cs, e := cc.GetChatService()
	if e != nil {
		utils.Println("[gateway] 聊天服务器为连接", e)
		return
	}
	_, e = cs.Logout(client.Id)
	if e != nil {
		utils.Println("[gateway] 用户退出转发至聊天服务器失败", e)
	}
	// 给区域内所有用户推送退出消息
	_, e = cs.ExitArea(client.Id, client.Area)
	if e != nil {
		utils.Println("[gateway] 用户退出区域转发至聊天服务器失败", e)
	}
}

// 切换区域/随机加入一个区域，没有就创建一个新的区域
func (gws *GatewayServer) ChangeArea(client *ws.Client, msg []byte) {
	utils.Println("[gateway] 用户随机加入区域", client.Id)
	ms, e := match.GetMatchService()
	if e != nil {
		return
	}
	cs, e := cc.GetChatService()
	if e != nil {
		return
	}
	// 给区域内所有用户推送退出消息
	_, _ = cs.ExitArea(client.Id, client.Area)
	// 退出游戏区域
	g := &game.Game{}
	// 随机加入到一个区域中
	res, e := ms.ChangeArea(client.Id)
	if e != nil {
		_ = client.Send(gws.Error("ChangeArea", e.Error()))
		return
	}
	// 退出游戏区域
	_, _ = g.DelArea(client.Id)
	// 设置新的游戏区域
	_, _ = g.SetArea(client.Id, res.Group, res.Area, res.GameId)
	// 给区域内所有用户推送加入通知
	_, _ = cs.JoinArea(client.Id, res.Area)
	client.Area = res.Area
	// 发送区域内在线用户列表
	list, e := ms.GetAreaInfo(client.Id)
	if e != nil {
		utils.Println(e)
		return
	}
	var ids []string
	for _, v := range list.Id {
		ids = append(ids, v.Id)
	}
	b, _ := json.Marshal(map[string]interface{}{
		"action": "AreaInfo",
		"list":   ids,
		"area":   list.Area,
		"game":   res.GameUrl,
		"gameId": res.GameId,
	})
	_ = client.Send(b)
}

// 发送消息
func (gws *GatewayServer) SendMsg(client *ws.Client, msg []byte) {
	type Req struct {
		ToId    string `json:"toId"` // 发给谁
		Content string // 消息内容
	}
	var data Req
	_ = json.Unmarshal(msg, &data)
	if data.ToId == "" {
		return
	}
	// 发送给聊天服务器转发给指定人
	utils.Println("[gateway] 用户消息 -> ", string(msg))
	// 给指定用户推送消息
	cs, e := cc.GetChatService()
	if e != nil {
		return
	}
	_, e = cs.SendMsg(client.Id, data.ToId, data.Content)
	if e != nil {
		utils.Println("[gateway] 用户消息发送失败 ->", e)
	}
}

// 区域内广播消息
func (gws *GatewayServer) SendMsgArea(client *ws.Client, msg []byte) {
	type Req struct {
		Content string // 消息内容
	}
	var data Req
	_ = json.Unmarshal(msg, &data)
	// 发送给聊天服务器转发给指定人
	utils.Println("[gateway] 用户广播消息 ->", client.Id, data.Content)
	// 给指定用户推送消息
	cs, e := cc.GetChatService()
	if e != nil {
		return
	}
	_, e = cs.SendMsgArea(client.Id, data.Content)
	if e != nil {
		utils.Println("[gateway] 用户广播消息发送失败 ->", e)
	}
}

// 获取区域信息
func (gws *GatewayServer) GetAreaInfo(client *ws.Client, msg []byte) {
	// 发送区域内在线用户列表
	ms, e := match.GetMatchService()
	if e != nil {
		utils.Println(e)
		return
	}
	list, e := ms.GetAreaInfo(client.Id)
	if e != nil {
		utils.Println(e)
		return
	}
	var ids []string
	for _, v := range list.Id {
		ids = append(ids, v.Id)
	}
	b, _ := json.Marshal(map[string]interface{}{
		"action": "AreaInfo",
		"list":   ids,
		"area":   list.Area,
		"game":   list.GameUrl,
		"gameId": list.GameId,
	})
	_ = client.Send(b)
}