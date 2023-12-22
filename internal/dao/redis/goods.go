package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mall/internal/constant"
	"mall/internal/entity"
	"mall/internal/logger"
	"time"
)

type GoodsRedisDao struct {
	redisCli *redis.Client
}

func NewGoodsRedisDao() GoodsRedisDao {
	return GoodsRedisDao{
		redisCli: GetRedisInstance(),
	}
}

func (dao GoodsRedisDao) GetGoodsInfo(ctx context.Context, goodsId int) (goods entity.GoodsInfo, err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsInfoPrefix, goodsId)
	result, err := dao.redisCli.Get(ctx, key).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Get err:%v", err)
		return goods, err
	}

	_ = json.Unmarshal([]byte(result), &goods)
	return
}

func (dao GoodsRedisDao) GetGoodsSkuInfo(ctx context.Context, skuId int) (goods entity.GoodsSkuInfo, err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsSkuPrefix, skuId)
	result, err := dao.redisCli.Get(ctx, key).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Get err:%v", err)
		return goods, err
	}

	_ = json.Unmarshal([]byte(result), &goods)
	return
}

func (dao GoodsRedisDao) SetGoodsSkuInfo(ctx context.Context, goods entity.GoodsInfo, expire time.Duration) (err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsSkuPrefix, goods.Id)
	ret, _ := json.Marshal(goods)
	_, err = dao.redisCli.Set(ctx, key, string(ret), expire).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Set err:%v", err)
		return err
	}

	return
}

func (dao GoodsRedisDao) SetGoodsInfo(ctx context.Context, goods entity.GoodsInfo, expire time.Duration) (err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsInfoPrefix, goods.Id)
	ret, _ := json.Marshal(goods)
	_, err = dao.redisCli.Set(ctx, key, string(ret), expire).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Set err:%v", err)
		return err
	}

	return
}

func (dao GoodsRedisDao) CheckSkuLeftStore(ctx context.Context, skuId int) bool {
	key := fmt.Sprintf(constant.RedisCacheGoodsSkuLeftStore, skuId)
	// decr ：减1；
	result, err := dao.redisCli.Decr(ctx, key).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Decr err:%v", err)
		return false
	}
	if result < 0 {
		return false
	}
	return true
}

func (dao GoodsRedisDao) IncrSkuLeftStore(ctx context.Context, skuId int) {
	key := fmt.Sprintf(constant.RedisCacheGoodsSkuLeftStore, skuId)
	_, err := dao.redisCli.Incr(ctx, key).Result()
	if err != nil {
		logger.WithContext(ctx).Errorf("redisClient.Incr err:%v", err)
		return
	}
	return
}
