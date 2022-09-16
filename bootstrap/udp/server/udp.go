package server

import (
	"encoding/json"
	"study-server/app/libs/encry"
	"study-server/app/libs/utils"
	"net"
)

type Message struct {
	Action string `json:"action"`
	Token  string `json:"token"`
}

// 客户端结构
type Client struct {
	server *UdpServer
	conn   *net.UDPAddr
	Id     string
}

// 开始监听,传入单次消息最大读取字节数，默认传1024即可
func (us *UdpServer) Star(size int) {
	for {
		var d = make([]byte, size)
		n, client, err := us.Server.ReadFromUDP(d[:]) // 读取客户端数据
		if err != nil {
			continue
		}
		go func(message []byte) {
			data := &Message{}
			e := json.Unmarshal(message, &data)
			// 非json数据不处理
			if e != nil {
				return
			}
			if data.Action == "Login" && data.Token != "" { // 登录到udp服务器
				info := encry.DecryptToken(data.Token)
				if info == nil {
					return
				}
				c := &Client{server: us, conn: client, Id: utils.ParamToString(info["id"])}
				us.clientList.Store(client.String(), c)
				us.userList.Store(utils.ParamToString(info["id"]), c)
				us.onLogin(c)
				return
			}
			res, ok := us.clientList.Load(client.String())
			if !ok { // 身份验证未通过
				_, _ = us.Server.WriteToUDP([]byte(`{"action":"Error","code":401,"msg":"401"}`), client)
				return
			}
			c := res.(*Client)
			if us.onUse != nil {
				us.onUse(c, &message)
			}
			if us.route[data.Action] != nil {
				us.route[data.Action](c, message)
			}
		}(d[:n])
	}
}

// 给当前客户端推送消息
func (c *Client) Send(msg []byte) error {
	_, e := c.server.Server.WriteToUDP(msg, c.conn)
	return e
}
