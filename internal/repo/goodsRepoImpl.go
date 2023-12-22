package repo

import (
	"context"
	"gorm.io/gorm"
	"mall/internal/dao/cache"
	"mall/internal/dao/db"
	"mall/internal/dao/redis"
	"mall/internal/entity"
	"time"
)

type GoodsRepositoryImpl struct {
	goodsDao      db.GoodsDbDao
	goodskuDao    db.GoodsSkuDbDao
	goodsRedisDao redis.GoodsRedisDao
	goodsCacheDao cache.GoodsBigCacheDao
}

func NewGoodsRepository() GoodsRepository {
	return &GoodsRepositoryImpl{
		goodsDao:      db.NewGoodsDbDao(),
		goodskuDao:    db.NewGoodsSkuDbDao(),
		goodsRedisDao: redis.NewGoodsRedisDao(),
		goodsCacheDao: cache.NewGoodsCacheDao(),
	}
}

func (repo *GoodsRepositoryImpl) WithDBInstance(db *gorm.DB) {
	repo.goodsDao.WithDBInstance(db)
	repo.goodskuDao.WithDBInstance(db)
	return
}

func (repo *GoodsRepositoryImpl) CreateGoods(ctx context.Context, goods entity.GoodsInfo) (goodsId int, err error) {
	goodsId, err = repo.goodsDao.CreateGoods(ctx, goods)
	if err != nil {
		return 0, err
	}

	// 更新关联id：sku.goodsid
	for idx, _ := range goods.SkuInfo {
		goods.SkuInfo[idx].GoodsId = goodsId
	}

	_, err = repo.goodskuDao.CreateGoodsSku(ctx, goods.SkuInfo)
	if err != nil {
		return 0, err
	}

	return
}

func (repo *GoodsRepositoryImpl) UpdateGoods(ctx context.Context, goods entity.GoodsInfo) (goodsId int, err error) {
	return repo.goodsDao.UpdateGoods(ctx, goods)
}

func (repo *GoodsRepositoryImpl) DeleteGoods(ctx context.Context, goods entity.GoodsInfo) (goodsId int, err error) {
	return repo.goodsDao.DeleteGoods(ctx, goods)
}

func (repo *GoodsRepositoryImpl) FindGoodsListByCategoryId(ctx context.Context, cateId int) (goodsList []entity.GoodsInfo, err error) {
	return repo.goodsDao.FindGoodsListByCategoryId(ctx, cateId)
}
func (repo *GoodsRepositoryImpl) GetGoodsDetailById(ctx context.Context, goodsId int) (goods entity.GoodsInfo, err error) {
	goods, _ = repo.goodsRedisDao.GetGoodsInfo(ctx, goodsId)
	if goods.Id > 0 {
		return
	}

	goods, err = repo.goodsDao.GetGoodsInfoById(ctx, goodsId)
	if err != nil {
		return goods, err
	}
	// 查询结果为空
	if goods.Id == 0 {
		return goods, nil
	}

	skuList, err := repo.goodskuDao.FindGoodsSkuByGoodId(ctx, goods.Id)
	if err != nil {
		return goods, err
	}
	goods.SkuInfo = skuList

	_ = repo.goodsRedisDao.SetGoodsInfo(ctx, goods, time.Second*60)
	return
}

func (repo *GoodsRepositoryImpl) SelectGoodsBySkuId(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {

	return repo.goodskuDao.SelectGoodsBySkuId(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) SelectGoodsBySkuIdForUpdate(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {
	return repo.goodskuDao.SelectGoodsBySkuIdForUpdate(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) UpdateGoodsSkuStore(ctx context.Context, sku entity.GoodsSkuInfo) (affectd int, err error) {
	return repo.goodskuDao.UpdateGoodsSkuStore(ctx, sku)
}

func (repo *GoodsRepositoryImpl) CheckSkuLeftStore(ctx context.Context, skuId int) bool {
	return repo.goodsRedisDao.CheckSkuLeftStore(ctx, skuId)
}

func (repo *GoodsRepositoryImpl) IncrSkuLeftStore(ctx context.Context, skuId int) {
	repo.goodsRedisDao.IncrSkuLeftStore(ctx, skuId)
}
