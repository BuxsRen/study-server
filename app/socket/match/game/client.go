package game

import (
	"errors"
	"study-server/app/libs/utils"
	gc "study-server/bootstrap/grpc/client"
	pg "study-server/bootstrap/grpc/proto/game"
	"sync"
)

type Game struct {
	server *pg.GameGrpcClient
	Url    string
	Id     string
}

var GameGrpc sync.Map // map[id]*Game
var GameArea sync.Map // map[areaid]*Game

// 网关服务器连接游戏grpc服务器
func (m *Game) Run(grpcUrl, id, udpUrl string) {
	utils.Println("[match][<=>][game] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	cgc := pg.NewGameGrpcClient(conn)
	GameGrpc.Store(id, &Game{server: &cgc, Url: udpUrl, Id: id})
}

// 获取该区域下分配的游戏服务器
func GetGameService(area string) (*Game, error) {
	var g *Game
	res, ok := GameArea.Load(area)
	if !ok {
		GameGrpc.Range(func(key, value interface{}) bool {
			g = value.(*Game)
			GameArea.Store(area, g)
			return true
		})
	} else {
		g = res.(*Game)
	}
	if g != nil {
		return g, nil
	}
	return nil, errors.New("匹配服务器 -> 未连接到 -> 游戏服务器")
}

// 随机获取一个在线的游戏服务器
func GetRandGameService() (*Game, error) {
	var g *Game
	GameGrpc.Range(func(key, value interface{}) bool {
		g = value.(*Game)
		return false
	})
	if g == nil {
		return nil, errors.New("没有在线的游戏服务器")
	}
	return g, nil
}