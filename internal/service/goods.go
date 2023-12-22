package service

import (
	"context"
	"mall/api/httputils"
	"mall/internal/dao/db"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/repo"
)

func CreateGoods(ctx context.Context, req entity.CreateGoodsReq) (resp entity.CreateGoodsResp, err error) {
	tx := db.GetDbInstance("").Begin()
	defer func() {
		if err != nil {
			logger.WithContext(ctx).Errorf("return error, rollback")
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	goodsRepo := repo.NewGoodsRepository()
	goodsRepo.WithDBInstance(tx)
	id, err := goodsRepo.CreateGoods(ctx, req.GoodsInfo)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.CreateGoods err:%v", err)
	}
	resp.GoodsId = id
	return
}

func UpdateGoods(ctx context.Context, req entity.UpdateGoodsReq) (resp entity.UpdateGoodsResp, err error) {
	goodsRepo := repo.NewGoodsRepository()
	_, err = goodsRepo.UpdateGoods(ctx, req.GoodsInfo)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.UpdateGoods err:%v", err)
	}
	return
}

func DeleteGoods(ctx context.Context, req entity.DeleteGoodsReq) (resp entity.DeleteGoodsResp, err error) {
	goodsRepo := repo.NewGoodsRepository()
	_, err = goodsRepo.DeleteGoods(ctx, req.GoodsInfo)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.DeleteGoods err:%v", err)
	}
	return
}

func GoodsList(ctx context.Context, req entity.GoodsListReq) (resp entity.GoodsListResp, err error) {
	goodsRepo := repo.NewGoodsRepository()
	goodsList, err := goodsRepo.FindGoodsListByCategoryId(ctx, req.CategoryId)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.FindGoodsListByCategoryId err:%v", err)
	}
	resp.GoodsList = goodsList
	return
}

func GetGoodsDetail(ctx context.Context, req entity.GetGoodsDetailReq) (resp entity.GetGoodsDetailResp, err error) {
	goodsRepo := repo.NewGoodsRepository()
	goods, err := goodsRepo.GetGoodsDetailById(ctx, req.GoodsId)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.GetGoodsDetailById err:%v", err)
	}
	if goods.Id == 0 {
		return resp, httputils.GoodsNotExists
	}

	resp.GoodsInfo = goods
	return
}
