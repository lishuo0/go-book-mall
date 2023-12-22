package constant

import "time"

// 分布式锁
const (
	RedisDistributeLockKeyPrefix = "redis_lock_%s"

	RedisDistributeLockExpire = 10 * time.Second
)

// 商品缓存
const (
	RedisCacheGoodsInfoPrefix = "redis_cache_goods_info_%d"

	RedisCacheGoodsSkuPrefix = "redis_cache_goods_sku_%d"

	RedisCacheGoodsInfoExpire = time.Second * 60 * 5

	RedisCacheGoodsSkuLeftStore = "cache_goods_sku_left_store_%d"
)
