package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2" // go get gopkg.in/yaml.v2
	"math/rand"
	"os"
	"study-server/bootstrap/helper"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	App = loadConfig()
}

// 加载 app.yaml 配置
func loadConfig() *app {
	path := flag.String("c", "./config/app.yaml", "输入 -c xxx.yaml 自定义配置文件")
	flag.Parse()
	var h = &helper.Helper{}
	file, e := os.ReadFile(*path)
	if e != nil {
		fmt.Println("➦ " + e.Error())
		h.Exit("✘ Config File ("+*path+")  Read Failed!", 3)
	}

	var app app
	e = yaml.Unmarshal(file, &app)
	if e != nil {
		fmt.Println("➦ " + e.Error())
		h.Exit("✘ Config Loading Failed!", 3)
	}
	fmt.Println("🔨 Config -> " + *path)
	return &app
}
