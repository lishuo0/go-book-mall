package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mall/internal/constant"
	"mall/internal/logger"
	"time"
)

type DistributeLockDao struct {
	redisCli *redis.Client
}

func NewDistributeLockDao() DistributeLockDao {
	return DistributeLockDao{
		redisCli: GetRedisInstance(),
	}
}

func (dao DistributeLockDao) Lock(ctx context.Context, suffix string, expire time.Duration) bool {
	key := fmt.Sprintf(constant.RedisDistributeLockKeyPrefix, suffix)
	result, err := dao.redisCli.SetNX(ctx, key, 1, expire).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.SetNX err:%v", err)
		return false
	}

	return result
}

func (dao DistributeLockDao) UnLock(ctx context.Context, suffix string) bool {
	key := fmt.Sprintf(constant.RedisDistributeLockKeyPrefix, suffix)
	result, err := dao.redisCli.Del(ctx, key).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Del err:%v", err)
		return false
	}

	return result == 1
}
