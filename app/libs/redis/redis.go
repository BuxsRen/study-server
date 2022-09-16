package redis

import (
	"context"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
	"study-server/bootstrap/redis"
	"time"
)

var ctx = context.Background()
var prefix = config.App.Redis.Prefix

// Redis 封装库
/**
 *@Example:
	rdb := redis.Redis{}
	fmt.Println(rdb.Get("test"))
*/
type Redis struct{}

type Api interface {
	Get(key string) string
	Set(key string, val interface{}) error
	Setex(key string, val interface{}, t int) error
	Del(key string) error
	RincrBy(key string, mode string) error
	HGet(name, field string) (string, error)
	HGetAll(name string) (map[string]string, error)
	HDel(name string, field ...string) error
	HLen(name string) error
	HExists(name, field string) (bool, error)
	HSet(name string, field ...interface{}) (bool, error)
	HSetNX(name, field string, value interface{}) (bool, error)
	CheckOut(name string)
}

// 取值。键
func (rdb *Redis) Get(key string) string {
	val, e := redis.GetPoolClinet().Get(ctx, prefix+key).Result()
	if e != nil {
		return ""
	}
	return val
}

// 存值。键，值
func (rdb *Redis) Set(key string, val interface{}) error {
	return redis.GetPoolClinet().Set(ctx, prefix+key, val, 0).Err()
}

// 存值，带过期时间，单位秒。键，值，秒
func (rdb *Redis) Setex(key string, val interface{}, t int) error {
	s := time.Duration(t) * time.Second
	return redis.GetPoolClinet().Set(ctx, prefix+key, val, s).Err()
}

// 删除值。键
func (rdb *Redis) Del(key string) error {
	return redis.GetPoolClinet().Del(ctx, prefix+key).Err()
}

// 递增或递减。键，模式(+,-)
func (rdb *Redis) RincrBy(key string, mode string) error {
	if mode == "+" {
		return redis.GetPoolClinet().IncrBy(ctx, prefix+key, 1).Err()
	} else {
		return redis.GetPoolClinet().IncrBy(ctx, prefix+key, -1).Err()
	}
}

// map 取值，map名称，map中的键名
func (rdb *Redis) HGet(name, field string) (string, error) {
	sc := redis.GetPoolClinet().HGet(ctx, prefix+name, field)
	return sc.Result()
}

// map 取全部
func (rdb *Redis) HGetAll(name string) (map[string]string, error) {
	return redis.GetPoolClinet().HGetAll(ctx, prefix+name).Result()
}

// map 删除map中对应的key
func (rdb *Redis) HDel(name string, field ...string) error {
	_, e := redis.GetPoolClinet().HDel(ctx, prefix+name, field...).Result()
	return e
}

// map 获取map的长度
func (rdb *Redis) HLen(name string) error {
	return redis.GetPoolClinet().HLen(ctx, prefix+name).Err()
}

// map 判断map中的指定键名是否存在
func (rdb *Redis) HExists(name, field string) (bool, error) {
	return redis.GetPoolClinet().HExists(ctx, prefix+name, field).Result()
}

// map 存值，map名称，map中的键名
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
func (rdb *Redis) HSet(name string, field ...interface{}) (bool, error) {
	return redis.GetPoolClinet().HMSet(ctx, prefix+name, field...).Result()
}

// map 存值，map名称，map中的键名，存在则保存失败，不存在则会创建
func (rdb *Redis) HSetNX(name, field string, value interface{}) (bool, error) {
	return redis.GetPoolClinet().HSetNX(ctx, prefix+name, field, value).Result()
}

// 输出保存的字符串缓存信息并退出,传入一个键名
func (rdb *Redis) CheckOut(name string) {
	str := rdb.Get(name)
	if str != "" {
		utils.ExitJson(str)
	}
}
