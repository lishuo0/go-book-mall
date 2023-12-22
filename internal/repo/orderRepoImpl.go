package repo

import (
	"context"
	"gorm.io/gorm"
	"mall/internal/dao/db"
	"mall/internal/entity"
	"strconv"
)

type OrderRepositoryImpl struct {
	orderDao       db.OrderDbDao
	orderDetailDao db.OrderDetailDbDao
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{
		orderDao:       db.NewOrderDbDao(),
		orderDetailDao: db.NewOrderDetailDbDao(),
	}
}

func (repo *OrderRepositoryImpl) WithDBInstance(db *gorm.DB) {
	repo.orderDetailDao.WithDBInstance(db)
	repo.orderDao.WithDBInstance(db)
	return
}

func (repo *OrderRepositoryImpl) CreateOrder(ctx context.Context, order entity.OrderInfo) (id int, err error) {
	id, err = repo.orderDao.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}
	for idx, _ := range order.OrderDetail {
		order.OrderDetail[idx].OrderId = strconv.Itoa(id)
	}

	err = repo.orderDetailDao.CreateOrderDetails(ctx, order.OrderDetail)
	if err != nil {
		return 0, err
	}

	return

}
func (repo *OrderRepositoryImpl) FindOrderList(ctx context.Context, userId int) ([]entity.OrderInfo, error) {
	return repo.orderDao.FindOrderList(ctx, userId)
}
func (repo *OrderRepositoryImpl) GetOrderDetail(ctx context.Context, orderId string) (order entity.OrderInfo, err error) {
	order, err = repo.orderDao.GetOrderInfo(ctx, orderId)
	if err != nil {
		return order, err
	}
	// 查询结果为空
	if order.Id == 0 {
		return order, nil
	}

	detailList, err := repo.orderDetailDao.FindOrderDetailList(ctx, orderId)
	if err != nil {
		return order, err
	}
	order.OrderDetail = detailList
	return
}
