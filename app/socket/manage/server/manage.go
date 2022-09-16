package server

import (
	"encoding/json"
	"study-server/app/libs/utils"
	"study-server/app/socket/manage/manage"
	pm "study-server/bootstrap/grpc/proto/manage"
	ws "study-server/bootstrap/websocket/server"
)

// 节点注册
func (ms *ManageServer) Login(client *ws.Client, force bool) {
	utils.Println("[manage] 节点注册", client.Id)
	// 给所有节点推送服务注册通知
	b, _ := json.Marshal(pm.List{
		Action: "ServerRegister",
		Node: &pm.Node{
			Id:   client.Id,
			Mode: client.Group,
			Rpc:  client.Rpc,
			Udp:  client.Udp,
		},
	})
	// 给所有的网关节点推送服务上线通知
	manage.WebSocket.SendAll(b, "")
}

// 节点注销
func (ms *ManageServer) Logout(client *ws.Client) {
	utils.Println("[manage] 节点注销", client.Id)
	b, _ := json.Marshal(pm.List{
		Action: "ServerLogout",
		Node: &pm.Node{
			Id:   client.Id,
			Mode: client.Group,
			Rpc:  client.Rpc,
			Udp:  client.Udp,
		},
	})
	// 给所有的网关节点推送服务离线通知
	manage.WebSocket.SendAll(b, "")
}