package server

import (
	"errors"
	"fmt"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
	"net"
	"strconv"
	"sync"
)

type UdpServer struct {
	Server     *net.UDPConn
	err        error
	route      map[string]func(*Client, []byte)
	onUse      func(client *Client, msg *[]byte) // 消息中间件/过滤器
	userList   sync.Map                          // 用户列表，通过客户端id取
	clientList sync.Map                          // 用户列表，通过客户端指针取
	onLogin    func(*Client)
}

type UdpServerInterface interface {
	Route() []*Route // 事件路由
	OnLogin(*Client)
}

// 事件路由
type Route struct {
	Action string                           // 路由名称或动作名称
	Fun    func(client *Client, msg []byte) // 对应处理的方法
}

// 初始化，传入事件路由
func New(route UdpServerInterface) *UdpServer {
	us := &UdpServer{}
	us.route = make(map[string]func(*Client, []byte))
	for _, v := range route.Route() {
		us.route[v.Action] = v.Fun
	}
	us.onLogin = route.OnLogin
	port, _ := strconv.Atoi(config.App.Server.Port)
	us.Server, us.err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	fmt.Printf("🌏 udp -> 127.0.0.1:%v\n", config.App.Server.Port)
	fmt.Printf("📡 Listen Server -> 0.0.0.0:%v\n", config.App.Server.Port)
	if us.err != nil {
		fmt.Println("➦ " + us.err.Error())
		(&helper.Helper{}).Exit("✘ This Port Is Already In Use !", 3)
		return nil
	}
	fmt.Printf("► OK! Start Service...\n\n")
	return us
}

// 给指定客户端推送消息
func (us *UdpServer) Send(id string, msg []byte) error {
	res, ok := us.userList.Load(id)
	if !ok {
		return errors.New("该客户端不存在")
	}
	c := res.(*Client)
	return c.Send(msg)
}
