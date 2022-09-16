package match

import (
	"errors"
	"study-server/app/libs/utils"
	gc "study-server/bootstrap/grpc/client"
	pm "study-server/bootstrap/grpc/proto/match"
)

type Match struct {
	service *pm.MatchGrpcClient
}

var MatchService *Match

// 聊天服务器连接匹配GRPC服务器
func (m *Match) Run(grpcUrl, id string) {
	utils.Println("[chat][<=>][match] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	c := pm.NewMatchGrpcClient(conn)
	MatchService = &Match{service: &c}
}

// 从已连接的匹配服务器中获取实例
func GetMatchService() (*Match, error) {
	if MatchService != nil {
		return MatchService, nil
	}
	return nil, errors.New("聊天服务器 -> 未连接到 -> 匹配服务器")
}