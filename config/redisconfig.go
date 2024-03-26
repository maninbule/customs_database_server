package config

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "172.21.116.147:6379",
		Password: "",
		DB:       0,
	})
	if Redis == nil {
		panic("Redis初始化失败")
	}
	//defer Redis.Close() 写了就直接关闭了
}
