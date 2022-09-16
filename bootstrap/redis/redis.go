package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"study-server/bootstrap/helper"
	"time"
)

var ctx = context.Background()

type Redis struct {
	Id        int           // 连接唯一id
	Client    *redis.Client // 连接实例
	heartbeat int           // 心跳时间(秒)
}

// 生成Redis连接实例,传入一个实例唯一Id
func NewRedis(id int) *Redis {
	r := &Redis{heartbeat: config.App.Redis.Ping, Id: id}
	r.Client = Dial()
	cmd := r.Client.Ping(ctx)
	e := cmd.Err()
	if e != nil {
		fmt.Println("➦ " + e.Error())
		(&helper.Helper{}).Exit("✘ Redis Connection Failed !", 3)
	}
	go r.ping()
	return r
}

// 连接Redis
func Dial() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.App.Redis.Host + ":" + config.App.Redis.Port,
		Password: config.App.Redis.PassWord,
		DB:       config.App.Redis.Db,
	})
}

// Redis 心跳
func (r *Redis) ping() {
	if r.heartbeat <= 0 {
		r.heartbeat = 60
	}
	for {
		time.Sleep(time.Duration(r.heartbeat+utils.Rand(30, 120)) * time.Second)
		r.Client.Ping(ctx)
	}
}
