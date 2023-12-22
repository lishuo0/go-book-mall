package redis

import (
	"github.com/go-redis/redis/v8"
	"mall/internal/core"
	"sync"
)

var rdb *redis.Client

var once sync.Once

func initRedis() {
	redisConfig := core.GlobalConfig.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password, // no password set
		DB:           redisConfig.Db,       // use default DB
		DialTimeout:  redisConfig.DialTimeout,
		ReadTimeout:  redisConfig.ReadTimeout,
		WriteTimeout: redisConfig.WriteTimeout,
	})
}

// 单例初始化
func GetRedisInstance() *redis.Client {
	if rdb != nil {
		return rdb
	}

	once.Do(func() {
		initRedis()
	})

	return rdb
}
