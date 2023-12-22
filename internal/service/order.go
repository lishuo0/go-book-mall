package service

import (
	"context"
	"github.com/spf13/cast"
	"mall/api/httputils"
	"mall/internal/dao/db"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/repo"
	"math/rand"
	"time"
)

// 假设只能买一个
func CreateOrderV2(ctx context.Context, req entity.CreateOrderReq) (resp entity.CreateOrderResp, err error) {
	// 模拟不同请求序列号，方式唯一索引冲突
	req.Serial = cast.ToString(time.Now().UnixNano()) + "_" + cast.ToString(rand.Intn(1000))

	goodsRepo := repo.NewGoodsRepository()
	if !goodsRepo.CheckSkuLeftStore(ctx, req.SkuId) {
		return resp, httputils.GoodsNotExists
	}

	job := NewOrderCreateJob(req)
	_ = NewAsynQueue(1000, 10).PushJob(ctx, job)
	// 异步处理去，不关注结果
	// 虽然客户端已经收到响应结果了，但是订单可能还没有创建成功
	return
}

func CreateOrderV1(ctx context.Context, req entity.CreateOrderReq) (resp entity.CreateOrderResp, err error) {
	// 模拟不同请求序列号，方式唯一索引冲突
	req.Serial = cast.ToString(time.Now().UnixNano()) + "_" + cast.ToString(rand.Intn(1000))
	goodsRepo := repo.NewGoodsRepository()
	if !goodsRepo.CheckSkuLeftStore(ctx, req.SkuId) {
		return resp, httputils.GoodsNotExists
	}

	tx := db.GetDbInstance("")
	// 返回时，如果出错了，回滚，否则提交事务
	defer func() {
		if err != nil {
			logger.WithContext(ctx).Errorf("return error, rollback")
			goodsRepo.IncrSkuLeftStore(ctx, req.SkuId)
			_ = tx.Rollback()
			return
		}

		_ = tx.Commit()
	}()

	// 开启事务
	tx = tx.Begin()

	// 查询sku详情
	goodsRepo.WithDBInstance(tx)
	sku, err := goodsRepo.SelectGoodsBySkuId(ctx, req.SkuId)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.SelectGoodsBySkuIdsForUpdate err:%v", err)
		return resp, httputils.CreateOrderFail
	}

	var order entity.OrderInfo
	order.TotalAmount = req.Count
	order.GoodsNum = req.Count
	order.Serial = req.Serial
	order.OrderDetail = append(order.OrderDetail, entity.OrderDetailInfo{
		SkuId: req.SkuId,
		Num:   req.Count,
		Price: sku.Price,
	})

	// 创建订单
	orderRepo := repo.NewOrderRepository()
	orderRepo.WithDBInstance(tx)
	_, err = orderRepo.CreateOrder(ctx, order)
	if err != nil {
		logger.WithContext(ctx).Errorf("orderRepo.CreateOrder err:%v", err)
		return resp, httputils.CreateOrderFail
	}
	resp.OrderId = ""
	return resp, nil
}

func CreateOrder(ctx context.Context, req entity.CreateOrderReq) (resp entity.CreateOrderResp, err error) {
	// 模拟不同请求序列号，方式唯一索引冲突
	req.Serial = cast.ToString(time.Now().UnixNano()) + "_" + cast.ToString(rand.Intn(1000))

	tx := db.GetDbInstance("")
	// 返回时，如果出错了，回滚，否则提交事务
	defer func() {
		if err != nil {
			logger.WithContext(ctx).Errorf("return error, rollback")
			_ = tx.Rollback()
			return
		}

		_ = tx.Commit()
	}()

	// 开启事务
	tx = tx.Begin()

	// 查询sku详情
	goodsRepo := repo.NewGoodsRepository()
	goodsRepo.WithDBInstance(tx)
	// for update：加锁，其他请求想更新商品库存，阻塞；（for share，加读锁）
	sku, err := goodsRepo.SelectGoodsBySkuIdForUpdate(ctx, req.SkuId)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.SelectGoodsBySkuIdForUpdate err:%v", err)
		return resp, httputils.CreateOrderFail
	}

	// 校验库存
	if req.Count > sku.Leftstore {
		return resp, httputils.GoodsNotEnough
	}

	var order entity.OrderInfo
	order.TotalAmount = req.Count
	order.GoodsNum = req.Count
	order.Serial = req.Serial
	order.OrderDetail = append(order.OrderDetail, entity.OrderDetailInfo{
		SkuId: req.SkuId,
		Num:   req.Count,
		Price: sku.Price,
	})

	sku.Leftstore -= req.Count
	_, err = goodsRepo.UpdateGoodsSkuStore(ctx, sku)
	if err != nil {
		logger.WithContext(ctx).Errorf("goodsRepo.UpdateGoodsSkuStore err:%v", err)
		return resp, httputils.CreateOrderFail
	}

	// 创建订单
	orderRepo := repo.NewOrderRepository()
	orderRepo.WithDBInstance(tx)
	_, err = orderRepo.CreateOrder(ctx, order)
	if err != nil {
		logger.WithContext(ctx).Errorf("orderRepo.CreateOrder err:%v", err)
		return resp, httputils.CreateOrderFail
	}
	resp.OrderId = ""
	return resp, nil
}

func OrderList(ctx context.Context, req entity.OrderListReq) (resp entity.OrderListResp, err error) {
	orderRepo := repo.NewOrderRepository()
	orderList, err := orderRepo.FindOrderList(ctx, req.UserId)
	if err != nil {
		logger.WithContext(ctx).Errorf("orderRepo.FindOrderList err:%v", err)
	}
	resp.OrderList = orderList
	return
}

func GetOrderDetail(ctx context.Context, req entity.GetOrderDetailReq) (resp entity.GetOrderDetailResp, err error) {
	orderRepo := repo.NewOrderRepository()
	order, err := orderRepo.GetOrderDetail(ctx, req.OrderId)
	if err != nil {
		logger.WithContext(ctx).Errorf("orderRepo.FindOrderList err:%v", err)
	}
	if order.Id == 0 {
		return resp, httputils.OrderNotExists
	}

	resp.OrderInfo = order
	return
}
