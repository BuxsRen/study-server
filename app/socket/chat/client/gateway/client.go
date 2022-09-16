package gateway

import (
	"errors"
	"study-server/app/libs/utils"
	gc "study-server/bootstrap/grpc/client"
	pg "study-server/bootstrap/grpc/proto/gateway"
	"sync"
)

type Gateway struct {
	service *pg.GatewayGrpcClient
}

var GatewayGrpc sync.Map // *Chat

// 聊天服务器连接网关GRPC服务器
func (m *Gateway) Run(grpcUrl, id string) {
	utils.Println("[chat][<=>][gateway] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	c := pg.NewGatewayGrpcClient(conn)
	gw := &Gateway{service: &c}
	GatewayGrpc.Store(id, gw)
}

// 从在线的网关列表中取一个实例
func GetGatewayService() (*Gateway, error) {
	var c *Gateway
	GatewayGrpc.Range(func(key, value interface{}) bool {
		c = value.(*Gateway)
		return true
	})
	if c != nil {
		return c, nil
	}
	return nil, errors.New("聊天服务器 -> 未连接到 -> 网关")
}

// 根据网关名称获取网关实例
func GetGateway(gateway string) (*Gateway, error) {
	res, ok := GatewayGrpc.Load(gateway)
	if ok {
		gw := res.(*Gateway)
		return gw, nil
	}
	return nil, errors.New("聊天服务器 -> 未连接到 -> 网关")
}

// 取所有网关
func GetAllGateway() []*Gateway {
	var list []*Gateway
	GatewayGrpc.Range(func(key, value interface{}) bool {
		list = append(list, value.(*Gateway))
		return true
	})
	return list
}