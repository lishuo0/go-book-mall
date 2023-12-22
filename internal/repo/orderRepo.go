package repo

import (
	"context"
	"gorm.io/gorm"
	"mall/internal/entity"
)

type OrderRepository interface {
	WithDBInstance(db *gorm.DB)
	CreateOrder(ctx context.Context, order entity.OrderInfo) (int, error)
	FindOrderList(ctx context.Context, userId int) ([]entity.OrderInfo, error)
	GetOrderDetail(ctx context.Context, orderId string) (entity.OrderInfo, error)
}
