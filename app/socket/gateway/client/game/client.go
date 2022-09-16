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

// 网关服务器连接游戏grpc服务器，传入grpc地址，游戏服务器id，udp地址
func (g *Game) Run(grpcUrl, id, udpUrl string) {
	utils.Println("[gateway][<=>][game] rpc ->", grpcUrl)
	conn := gc.Dial(grpcUrl)
	cgc := pg.NewGameGrpcClient(conn)
	GameGrpc.Store(id, &Game{server: &cgc, Url: udpUrl, Id: id})
}

// 通过游戏服务器id获取游戏服务器实例
func (g *Game) GetGame(id string) (*Game, error) {
	res, ok := GameGrpc.Load(id)
	if !ok {
		return nil, errors.New("没有这个服务器，或者未连接")
	}
	return res.(*Game), nil
}

func (g *Game) GetAllService() []*Game {
	var list []*Game
	GameGrpc.Range(func(key, value interface{}) bool {
		gs := value.(*Game)
		list = append(list, gs)
		return true
	})
	return list
}