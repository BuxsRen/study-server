# 一个边学习Go边写的个人项目

2022-01 / 2022-06

 - 练习Go基础语法，管道，map等
 - 练习Gin框架与ORM
 - 练习WebSocket
 - 练习Grpc与protobuf
 - php转的go，还有点php的影子，目录根据laravel来的

> 现在转goframe框架了，前面一知半解写的东西，很烂，放git上做个记录，日后看到也好刺激自己，让自己奋发图强，写出更好的代码来

#### 主要有哪些功能？
1. Web框架 MVC分层，Laravel风格，路由，控制器，中间件，视图，模型，仿Laravel参数验证
2. WebSocket框架，已实现聊天，私聊，群聊，上下线通知，被迫下线通知等
3. 聊天区域划分，登录随机分配到空闲的区域，区域人数爆满会自动创建一个新的区域，也有私人区域，别人无法通过匹配加入
4. 集成大量第三方库，短信，文件上传，邮件，钉钉/企业微信推送，支付，二维码，验证码，加解密，数据xls导出等

#### 如何玩转？
1. 启动管理服务端
2. 启动一个webserver，前端账号密码登录拿到token
3. 然后分别启动聊天服务，匹配服务
4. 启动一个网关(可以启动多个，多个网关之间可以互相聊天)，前端ws接入，传入token，登录到网关
5. 游戏服务是udp发送需要实时的数据的，例如语音通话功能(详见Hello App,仿微信,React Native + Java))

#### 相关目录说明
```text
├──app # 主要功能逻辑实现区
│   ├──console # 后台任务，触发事件
│   ├──events # 事件注册
│   ├──http # http服务
│   ├──libs # 通用库
│   ├──models # 数据库模型
│   │    └──models.go # 数据库模型基类
│   │
│   └──socket
│       ├──chat # 聊天服务
│       ├──gateway # 网关服务
│       ├──match # 匹配服务
│       ├──game # 游戏服务
│       ├──manage # 管理服务
│       └──middleware # ws中间件
│
├──bootstrap # 启动，初始化
│   ├──area # 区域(房间)分配主要实现逻辑
│   ├──config # 配置文件加载
│   ├──database # 数据库连接与日志
│   ├──exceptions   # 异常处理
│   ├──grpc # grpc服务注册
│   ├──helper # 助手类
│   ├──redis # redis连接
│   ├──request # 请求参数收集
│   ├──router # 路由注册
│   ├──server # 服务启动
│   ├──udp # udp服务注册
│   └──websocket # ws服务注册
│
├──config  # 配置文件所在目录
├──protobuf  # protobuf文件
├──resource  # 模板文件
├──routes # 用户路由
├──storage # 日志

```

#### 环境与包
| Golang | MySQL | Nginx | Redis |  Gin  |  Gorose  |
|:------:|:-----:| :---: | :---: | :---: | :------: |
|  1.18  |  5.7  | 1.18  |  6.2  |  1.7  |  2.1.5   |

#### 安装依赖

-  安装 protoc
> https://github0.com/protocolbuffers/protobuf/releases

```shell
# 启用 go mod
go env -w GO111MODULE=on
#使用七牛云代理
go env -w GOPROXY=https://goproxy.cn,direct

# go mod init study-server
# go mod tidy

# 安装 protoc go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

#### 前置

- 复制 app.yaml.example 内容 到 app.yaml

#### 配置 /config/app.yaml 配置文件
- 根据需要进行配置，用不到的配置不用配
```yaml
server: # *服务配置
  mode: gateway # 服务模式 manage(管理服务器[单])、webserver(接口服务器[多])、gateway(网关服务器[多])、chat(聊天服务器[多])、match(匹配服务器[单])、game(游戏服务器[多])
  url: http://127.0.0.1:9310 # 项目地址/域名/项目访问地址/上传资源访问地址
  host: 127.0.0.1 # 服务监听地址，推荐 127.0.0.1
  port: 6001 # 服务监听端口
  grpc_port: 6002 # grpc服务监听端口
  ip: 192.168.1.100 # 服务器ip，内网或者公网ip
  debug: false # 开启debug模式
  env: local # 运行环境 local(开发) production(线上)
  log_access: ./storage/logs/go_access.log # 访问日志保存路径
  log_error: ./storage/logs/go_error.log # 错误日志保存路径
  template: false # 加载模板 false 的时候 部署可以不需要resources目录
  manage_ws: ws://192.168.1.99:6001/ws/manage # 管理服务器ws地址，当服务模式不是manag的时都需要提供这个地址，这个地址是manage服务器的地址，只能运行一个manage
  manage_tcp: 192.168.1.99:6002 # 管理服务器grpc地址

mysql: # *mysql 配置
  host: 127.0.0.1  # 数据库地址
  port: 3306       # 数据库端口
  database: xxxxx  # 数据库名
  username: root   # 用户名
  password: 123456 # 密码
  prefix: unite_     # 表前缀
  log: false        # 开启sql日志，打印sql执行日志到控制台(server.debug模式打开的时候才会输出到控制台)
  save_log: false   # 保存sql日志到文本，需要先打开 "开启sql日志"，server.debug模式关闭的时候照样可以写入到文件
  log_path: ./storage/logs/sql.log # sql 日志保存地址，需要先打开 "保存sql日志到文本"

redis: # *redis 配置
  host: 127.0.0.1  # 地址
  port: 6379 # 端口`
  password: # 密码
  prefix: unite_ # redis前缀
  pool: 3 # 连接池数量
  ping: 60 # 心跳时间
  db: 0 # 使用几号数据库
  
rsa: # *RSA 加解密
  public_key: ./config/rsa_public.key # 公钥路径
  private_key: ./config/rsa_private.key # 私钥路径

other: # *其他配置
  public_dir: /www/wwwroot/public/upload/ # 静态文件保存目录，后面一定要加上 / ,其中 /www/wwwroot 是nginx的静态资源目录
  public_prefix: /public/upload/ # 前端寻址前缀
  token_key: xxxxxxxxxxxx # 接口token签发密钥
  aes_key:  xxxxxxxxxxxxxxxx # AES 16位加密字符串密码

push: # 推送配置
  use: false # 开启异常推送(env是production时可用)
  mode: bark # 推送方式：bark、dingTalk、dingTalkMarkDown、wechat
  bark_url: https://api.day.app/xxxxx/ # bark推送地址
  dingTalk_url: https://oapi.dingtalk.com/robot/send?access_token=xxxxxxx # 钉钉推送地址
  wechat_url: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx # 企业微信推送地址

qiniuyun: # 七牛云配置
  bucket: xxx # 七牛云空间名称
  access_key: xxxxxxxxxxx # 七牛云AK
  secret_key: xxxxxxxxxxx # 七牛云SK

email: # 邮箱配置
  name: Break技术团队 # 发件人名称
  user: xxxx@163.com # 发件人邮箱
  pass: xxxxxxxxxxxx # 发件人密码
  host: smtp.163.com # 邮箱服务器
  port: 465 # 邮箱端口

alipay: # 支付宝网页&移动应用支付配置
  appid: 20220122224500 # AppID
  private_key: ./config/alipay.key # 应用私钥路径，后缀也可以是txt，反正是文本就行
  notify_url: https://www.xxx.com/xxxx # 回调地址
```

#### 编译运行
```shell script
# 编译
go build main.go
# 运行
./main -c /home/app.yaml 
# 或者直接 
./main

# Windows 环境下使用build.bat 一键生成打包文件
```

#### 配置nginx反向代理
```shell script
# http
location / {
    proxy_set_header X-Forward-For $remote_addr;
    proxy_set_header X-real-ip $remote_addr;
    proxy_set_header Host $http_host;
    proxy_set_header SERVER_PORT $server_port;
    proxy_set_header REMOTE_ADDR $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://ip:port;
}

# WebSocket
location /ws {
    proxy_pass http://ip:port;
    proxy_set_header Host $host:$server_port;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "Upgrade";
}
```