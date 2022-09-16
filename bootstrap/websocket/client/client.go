package client

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

type WebSocket struct {
	deviceNumber string                              // 设备唯一编号
	conn         *websocket.Conn                     // ws连接
	message      chan []byte                         // 发送消息通道
	status       bool                                // ws状态
	token        string                              // 登录token 通过登录接口获取
	url          string                              // 登录地址
	onUse        func(*WebSocket, *[]byte)           // 消息中间件/过滤器
	route        map[string]func(*WebSocket, []byte) // 路由列表
	i            WebSocketInterface
}

// 必须实现下面四个方法
type WebSocketInterface interface {
	Route() []*Route       // 事件路由
	OnOpen(ws *WebSocket)  // 连接成功回调事件
	OnClose(ws *WebSocket) // 断开连接回调事件
	OnError(ws *WebSocket) // 连接失败回调事件
}

// 事件路由
type Route struct {
	Action string                          // 路由名称或动作名称
	Fun    func(ws *WebSocket, msg []byte) // 对应处理的方法
}

type Message struct {
	Action  string `json:"action"`
	Content string `json:"content"`
	FromId  string `json:"fromId"`
}

// 初始化,传入设备唯一id和登录地址 ws://ip:port?xxx=xxxxxx
func New(id, url string) *WebSocket {
	ws := &WebSocket{deviceNumber: id, url: url}
	ws.route = make(map[string]func(*WebSocket, []byte))
	return ws
}

// 绑定事件路由
func (ws *WebSocket) Bind(i WebSocketInterface) {
	for _, v := range i.Route() {
		ws.route[v.Action] = v.Fun
	}
	ws.i = i
}

// 是否断开重连 , 重连间隔时间(秒) 启动后会阻塞在该方法处
func (ws *WebSocket) Start(again bool, timeout int) {
	for {
		if ws.status == false {
			ws.dial()
		}
		time.Sleep(time.Second * time.Duration(timeout))
		if !again {
			break
		}
	}
}

// 连接服务端
func (ws *WebSocket) dial() {
	var err error
	ws.conn, _, err = websocket.DefaultDialer.Dial(ws.url, nil) // 连接服务端
	if err != nil {
		ws.i.OnError(ws)
		return
	}
	ws.message = make(chan []byte, 256)
	ws.status = true
	ws.i.OnOpen(ws)
	go ws.onSend()
	ws.onMessage()
}

// 读取消息
func (ws *WebSocket) onMessage() {
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil { // 出现错误，退出读取，尝试重连
			ws.i.OnClose(ws)
			_ = ws.conn.Close()
			ws.conn = nil
			time.Sleep(1 * time.Second)
			if ws.message != nil {
				close(ws.message)
				ws.message = nil
			}
			ws.status = false
			break
		}
		if len(message) == 0 {
			continue
		}
		var data Message
		e := json.Unmarshal(message, &data)
		if e != nil {
			continue
		}
		if data.Action == "Ping" {
			continue
		}
		if ws.onUse != nil {
			ws.onUse(ws, &message)
		}
		if ws.route[data.Action] != nil {
			go ws.route[data.Action](ws, message)
		}
	}
}

// 发送消息
func (this *WebSocket) onSend() {
	for {
		select {
		case msg, ok := <-this.message:
			if !ok {
				return
			}
			_ = this.conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// 消息拦截器/中间件，消息到达控制器前先经过中间件
func (ws *WebSocket) Use(f func(ws *WebSocket, msg *[]byte)) {
	ws.onUse = f
}

// 给服务端推送消息
func (ws *WebSocket) Send(msg []byte) error {
	if ws != nil && ws.conn != nil && ws.message != nil {
		ws.message <- msg
		return nil
	} else {
		return errors.New("消息失败,连接不可用")
	}
}
