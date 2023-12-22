package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"mall/internal/constant"
	"mall/internal/entity"
	"mall/internal/logger"
)

type GoodsBigCacheDao struct {
	bigcacheCli *bigcache.BigCache
}

func NewGoodsCacheDao() GoodsBigCacheDao {
	return GoodsBigCacheDao{
		bigcacheCli: GetBigCacheInstance(),
	}
}

func (dao GoodsBigCacheDao) GetGoodsInfo(ctx context.Context, goodsId int) (goods entity.GoodsInfo, err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsInfoPrefix, goodsId)
	result, err := dao.bigcacheCli.Get(key)
	if err != nil {
		logger.WithContext(ctx).Errorf("bigcacheCli.Get err:%v", err)
		return goods, err
	}
	_ = json.Unmarshal(result, &goods)
	return
}

func (dao GoodsBigCacheDao) SetGoodsInfo(ctx context.Context, goods entity.GoodsInfo) (err error) {
	key := fmt.Sprintf(constant.RedisCacheGoodsInfoPrefix, goods.Id)
	ret, _ := json.Marshal(goods)
	err = dao.bigcacheCli.Set(key, ret)
	if err != nil {
		logger.WithContext(ctx).Errorf("bigcacheCli.Set err:%v", err)
		return err
	}

	return
}
