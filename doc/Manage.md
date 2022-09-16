# 管理服务端开发手册

#### 管理端代码位置
##### app/socket/manage

```text
├──grpc  # Grpc 服务端模块   			
│	└──manage.go  # 提供grpc服务给客户端调用
│
├──manage # 一些参数，结构等初始化、配置等 
│	├──manage.go      
│
├──server  # Websocket 服务端模块
│	├──manage.go  # 控制器
│	├──route.go  # 事件路由，用于绑定到控制器
│	└──server.go  # 客户端接入WebSocket，并开始监听客户端的消息和解析消息
```
---
#### protobuf 文件地址 
##### protobuf/manage.proto

---
#### proto 生成的 GRPC 文件地址
##### bootstarp/grpc/proto/manage/*

---
# 快速开始

##### 连接到管理端(注册服务)

> 编辑配置文件配置正确的ip和端口

#### 配置连接信息
```yaml
server:
  manage_ws: ws://ip:port/ws/manage
  manage_tcp: ip:port
```

#### 服务连接并注册到管理端
```go
// 生成token
token := encry.EncryptToken("需要注册的服务的唯一Id",-1,map[string]interface{}{
    "group": config.App.Server.Mode, // 注册的组，传当前运行的模式
    "ws": "ws://ip:port/ws/xxx",     // 注册的ws地址，传当前服务可以正常连接到的ws地址
    "tcp": "ip:port",  // 注册的grpc地址，如果当前服务使用到grpc，需要传这个
})
// 生成管理端 websocket 连接地址，
wsurl := config.App.Server.ManageWs + "?token="+token

// 连接到管理端WebSocket，自动注册服务
websocket.Dial(wsurl)
// 连接到管理端GRPC服务以获得管理端GRPC能力
grpc.Dial(config.App.Server.ManageWs,grpc.WithInsecure())

```

#### 管理端能力

##### 1.获取服务(grpc)
```
// 获取聊天服务器列表，返回聊天服务器的ws、tcp连接地址
func GetService(context.Background(), &pm.Node{Mode: "chat"})
```

##### 2.服务上线通知(websocket)
- 该函数绑定到WebSocket客户端的事件路由上
```
// 服务注册通知，返回服务的类型以及ws、tcp连接地址
func ServerRegister(socket *ws.WebSocket,msg []byte)
```