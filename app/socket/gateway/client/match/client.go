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

// 网关服务器连接匹配grpc服务器
func (m *Match) Run(grpcUrl, id string) {
	utils.Println("[gateway][<=>][match] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	c := pm.NewMatchGrpcClient(conn)
	MatchService = &Match{service: &c}
}

// 从在线的匹配服务器中取一个连接
func GetMatchService() (*Match, error) {
	if MatchService != nil {
		return MatchService, nil
	}
	return nil, errors.New("网关 -> 未连接到 -> 匹配服务器")
}