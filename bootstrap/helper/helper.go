package helper

import (
	"fmt"
	"log"
	"os"
	"time"
)

var LOG *log.Logger // 日志

// 助手类
type Helper struct {}

// 输出信息并在3秒后退出
func (this *Helper) Exit(message string,s int){
	fmt.Println(message)
	fmt.Printf("✘ Error, Service Running Error!!! \n\n")
	time.Sleep(time.Duration(s) * time.Second)
	os.Exit(0)
}

// 取 .env 中的值
func (this *Helper) GetEnv(name string) string {
	return os.Getenv(name)
}

// 模板视图
type Views struct {
	Template string
	Data interface{}
}

// 控制器输出视图 return helper.View("404.tpl",nil)
func View(template string, data interface{}) Views {
	return Views{
		Template: template,
		Data: data,
	}
}


// 写日志，保存在错误日志文件中，提供给未开启debug模式时使用
func Log(str string) {
	if LOG != nil {
		_ = LOG.Output(2, "\n"+str+"\n\n")
	}
}
