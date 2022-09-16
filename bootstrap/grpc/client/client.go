package client

import (
	"fmt"
	"google.golang.org/grpc"
)

func Dial(url string) *grpc.ClientConn {
	conn,e := grpc.Dial(url,grpc.WithInsecure()) // 这个错误好像不会抛
	if e != nil {
		fmt.Println("grpc 服务连接失败 ->",e)
	}
	return conn
}