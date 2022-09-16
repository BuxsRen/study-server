package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

// Redis连接池
var redisPools *pool

type pool struct {
	Pools []*Redis
}

func init() {
	redisPools = newPool()
}

// 初始化连接池
func newPool() *pool {
	rp := &pool{}
	if config.App.Redis.Pool <= 0 {
		config.App.Redis.Pool = 3
	}
	for i := 1; i <= config.App.Redis.Pool; i++ {
		rp.Pools = append(rp.Pools, NewRedis(i))
	}
	fmt.Printf("❤ Redis -> [%d] %v:%v\n", config.App.Redis.Pool, config.App.Redis.Host, config.App.Redis.Port)
	return rp
}

// 连接池中随机获取一个链接
func GetPoolClinet() *redis.Client {
	index := utils.Rand(0, len(redisPools.Pools))
	return redisPools.Pools[index].Client
}
