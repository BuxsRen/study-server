package server

import (
	"fmt"
	"github.com/gin-gonic/gin" // https://gin-gonic.com/zh-cn/docs/ Gin开发文档
	"io"
	"log"
	"os"
	"runtime"
	"study-server/app/libs/utils"
	"study-server/app/socket/chat/client/chat"
	cs "study-server/app/socket/chat/grpc"
	gc "study-server/app/socket/game/client"
	gp "study-server/app/socket/game/grpc"
	gs "study-server/app/socket/game/server"
	gwc "study-server/app/socket/gateway/client/client"
	gws "study-server/app/socket/gateway/grpc"
	gm "study-server/app/socket/manage/grpc"
	mc "study-server/app/socket/match/client"
	mg "study-server/app/socket/match/grpc"
	"study-server/bootstrap/config"
	"study-server/bootstrap/exceptions"
	"study-server/bootstrap/helper"
	"study-server/bootstrap/routes"
)

type server struct {
	server    *gin.Engine
	h         *helper.Helper
	host      string
	port      string
	debug     bool
	logAccess string
	LogError  string
	template  bool
}

// 启动应用
func Run() {
	app := &server{
		host:      config.App.Server.Host,
		port:      config.App.Server.Port,
		debug:     config.App.Server.Debug,
		logAccess: config.App.Server.LogAccess,
		LogError:  config.App.Server.LogError,
		template:  config.App.Server.Template,
	}
	app.print()
	app.h = &helper.Helper{}
	app.isDebug()
	var req = "ws://"
	switch config.App.Server.Mode {
	case "game":
		go gc.Run()
		go gp.Star()
		gs.Star()
	case "match":
		go mg.Star()
		go mc.Run()
	case "manage":
		go gm.Star()
	case "gateway":
		go gwc.Run()
		go gws.Star()
	case "chat":
		go chat.Run()
		go cs.Star()
	default:
		req = "http://"
	}
	ip := "127.0.0.1"
	if config.App.Server.Host != "127.0.0.1" {
		ip = config.App.Server.Ip
	} else if config.App.Server.Host != "0.0.0.0" {
		ip = config.App.Server.Host
	}
	fmt.Printf("🌏 %s%s:%s\n", req, ip, app.port)
	fmt.Printf("🔗 Listen Server -> %s:%s\n", ip, app.port)
	fmt.Printf("► OK! Start Service ...\n")
	app.server = gin.Default()
	_ = app.server.SetTrustedProxies(nil)
	app.loadTemplate()
	app.server.Use(exceptions.Handle)            // 异常处理
	(&routes.Route{Router: app.server}).Handle() // 加载路由
	app.start()
}

// 输出信息
func (s *server) print() {
	fmt.Printf("🌙 Mode -> %s\n", config.App.Server.Mode)
	fmt.Printf("⏰ Runtime: %s\n", utils.GetNow().Format("2006-01-02 15:04:05"))
	fmt.Printf("🚀 Server: 2021 - %d By Break\n", utils.GetNow().Year())
	fmt.Printf("💻 System：%s (%s)\n", runtime.GOOS, runtime.GOARCH)
}

// 是否开启debug
func (s *server) isDebug() {
	if !s.debug { //未开启debug 异常错误日志写入文件中,开启debug 异常打印到终端并返回给浏览器
		gin.SetMode(gin.ReleaseMode)
		access, _ := os.Create(s.logAccess)
		gin.DefaultWriter = io.MultiWriter(access) // 访问日志
		s.logs()
	}
}

// 是否加载模板
func (s *server) loadTemplate() {
	if s.template { // 加载模板
		s.server.LoadHTMLGlob("./resources/views/**/*")
	}
}

// 创建输出日志文件
func (s *server) logs() {
	logFile, e := os.Create(config.App.Server.LogError)
	if e != nil {
		fmt.Println("➦ " + e.Error())
		s.h.Exit("✘ Error Log File Create Failed !", 3)
	}
	helper.LOG = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// 开始web服务
func (s *server) start() {
	err := s.server.Run(s.host + ":" + s.port) // 启动服务
	if err != nil {
		fmt.Println("➦ " + err.Error())
		s.h.Exit("x s Port Is Already In Use!", 0)
	}
}
