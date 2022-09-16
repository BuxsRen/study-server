package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2" // https://github.com/gohouse/gorose & https://www.kancloud.cn/fizz/gorose/769179
	"log"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
	"time"
)

type info struct {
	user     string
	pass     string
	host     string
	port     string
	database string
	prefix   string
}

type Mysql struct {
	Id        int           // 唯一Id
	Conn      *gorose.Engin // 连接实例
	heartbeat int           // 心跳时间(秒)
	log       *log.Logger   // 日志
	info      *info
}

// 生成Mysql连接实例,传入一个实例唯一Id，日志文件句柄
func NewMysql(id int) *Mysql {
	my := &Mysql{Id: id, heartbeat: config.App.Mysql.Ping}
	my.info = getConfig()
	conn, e := Dial(my.info)
	if e != nil {
		fmt.Println("➦ " + e.Error())
		(&helper.Helper{}).Exit("✘ Mysql Connection Failed !", 3)
	}
	my.Conn = conn
	if config.App.Mysql.Log { // 数据库日志
		my.Conn.Use(func(eg *gorose.Engin) {
			eg.SetLogger(NewLogger(&LogOption{
				EnableSqlLog:   true, // sql日志
				EnableSlowLog:  5,
				EnableErrorLog: true, // 错误日志
			}))
		})
	}
	my.ping()
	return my
}

func (my *Mysql) ping() {
	if my.heartbeat <= 0 {
		my.heartbeat = 60
	}
	go func() {
		for {
			time.Sleep(time.Duration(my.heartbeat+utils.Rand(30, 120)) * time.Second)
			_ = my.Conn.Ping()
		}
	}()
}

// 连接mysql
func Dial(con *info) (*gorose.Engin, error) {
	return gorose.Open(&gorose.Config{
		Driver: "mysql",
		Dsn:    fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=true", con.user, con.pass, con.host, con.port, con.database),
		Prefix: con.prefix,
	})
}

func getConfig() *info {
	return &info{
		user:     config.App.Mysql.UserName,
		pass:     config.App.Mysql.PassWord,
		host:     config.App.Mysql.Host,
		port:     config.App.Mysql.Port,
		database: config.App.Mysql.Database,
		prefix:   config.App.Mysql.Prefix,
	}
}
