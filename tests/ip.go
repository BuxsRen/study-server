package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"study-server/app/libs/net"
)

type Ip struct{}

func main() {
	ip := &Ip{}
	fmt.Println(ip.GetNetworkIp())
	fmt.Println(ip.GetLocalIP())
}

func (ip *Ip) GetNetworkIp() string {
	res, e := net.New("https://www.cip.cc/", "GET", "").Do()
	if e != nil {
		fmt.Println(e)
		return ""
	}
	str := ip.regexpStr(res.(string), `IP	: \d+.\d+.\d+.\d+`)
	if str != "" {
		str = ip.regexpStr(str, `\d+.\d+.\d+.\d+`)
	}
	return str
}

func (ip *Ip) GetLocalIP() string {
	var str string
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd.exe", "/c", "chcp 65001 & ipconfig")
		b, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return ""
		}
		str2 := ip.regexpStr(string(b), `IPv4 Address. . . . . . . . . . . : \d+.\d+.\d+.\d+`)
		if str2 != "" {
			str = ip.regexpStr(str2, `\d+.\d+.\d+.\d+`)
		}
	case "linux":
		cmd := exec.Command("/bin/sh", "-c", "ifconfig")
		b, err := cmd.Output()
		if err != nil {
			cmd = exec.Command("/bin/sh", "-c", "ip addr")
			b, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return ""
			}
			str2 := ip.regexpStr(string(b), `inet \d+.\d+.\d+.\d+/24 brd`)
			if str2 != "" {
				str = ip.regexpStr(str2, `\d+.\d+.\d+.\d+`)
			}
			return str
		}
		str2 := ip.regexpStr(string(b), `inet \d+.\d+.\d+.\d+  netmask`)
		if str2 != "" {
			str = ip.regexpStr(str2, `\d+.\d+.\d+.\d+`)
		}
	}
	return str
}

func (ip *Ip) regexpStr(str, reg string) string {
	word := regexp.MustCompile(reg).FindAllStringSubmatch(str, 1)
	if len(word) > 0 && len(word[0]) > 0 {
		return word[0][0]
	}
	return ""
}
