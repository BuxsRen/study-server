package database

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"log"
	"os"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
)

// Redis连接池
var mysqlPool *pool
var LogFile *log.Logger

type pool struct {
	Pools []*Mysql
}

func init() {
	mysqlPool = newPool()
}

// 初始化连接池
func newPool() *pool {
	rp := &pool{}
	if config.App.Mysql.Pool <= 0 {
		config.App.Mysql.Pool = 3
	}
	file, e := os.Create(config.App.Mysql.LogPath)
	if e != nil {
		fmt.Println("➦ " + e.Error())
		(&helper.Helper{}).Exit("✘ Mysql Log File Create Failed !", 3)
	}
	if config.App.Mysql.Log && config.App.Mysql.SaveLog {
		LogFile = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	for i := 1; i <= config.App.Mysql.Pool; i++ {
		rp.Pools = append(rp.Pools, NewMysql(i))
	}
	fmt.Printf("🐬 Mysql -> [%d] @tcp(%v:%v)/%v\n", config.App.Mysql.Pool, rp.Pools[0].info.host, rp.Pools[0].info.port, rp.Pools[0].info.database)
	return rp
}

// 连接池中随机获取一个链接
func GetPoolClinet() *gorose.Engin {
	index := utils.Rand(0, len(mysqlPool.Pools))
	return mysqlPool.Pools[index].Conn
}
