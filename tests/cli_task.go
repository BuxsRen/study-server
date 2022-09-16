package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "测试命令行程序"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "cmd:name",               // 命令的名字
			Aliases: []string{"t:s"},          // 命令的缩写
			Usage:   "运行任务 命令： cmd:name xxxx", // 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Action: func(c *cli.Context) error { // 命令的处理函数
				run(c.Args().Get(0))
				return nil
			},
		},
	}
	_ = app.Run(os.Args)
}

func run(name string) {
	fmt.Println(name)
}
