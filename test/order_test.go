package test

import (
	"context"
	"github.com/hashicorp/go-uuid"
	"gorm.io/sharding"
	"mall/internal/core"
	"mall/internal/dao/db"
	"mall/internal/entity"
	"mall/internal/logger"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	_ = core.InitConfig("/Users/lile/Documents/gocode/go-class/mall/configs/conf.yaml")
	_ = logger.InitLogger()

	uuids, _ := uuid.GenerateUUID()
	tx := db.GetDbInstance("")
	_ = tx.Use(sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      4,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "mall_order"))
	orderDao := db.NewOrderDbDao().WithDBInstance(tx)
	_, err := orderDao.CreateOrder(context.Background(), entity.OrderInfo{
		OrderId:     uuids,
		UserId:      125404,
		TotalAmount: 100,
		GoodsNum:    1,
		Serial:      uuids,
	})
	if err != nil {
		t.Errorf("create order err:%v", err)
	}

	time.Sleep(time.Second * 1)

}

func TestGetOrderInfo(t *testing.T) {
	_ = core.InitConfig("/Users/lile/Documents/gocode/go-class/mall/configs/conf.yaml")
	_ = logger.InitLogger()

	tx := db.GetDbInstance("")
	_ = tx.Use(sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      4,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "mall_order"))
	orderDao := db.NewOrderDbDao().WithDBInstance(tx)
	_, err := orderDao.FindOrderList(context.Background(), 125404)
	if err != nil {
		t.Errorf("get order err:%v", err)
	}

	time.Sleep(time.Second * 1)

}
