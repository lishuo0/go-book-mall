package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mall/internal/entity"
	"mall/internal/logger"
	"time"
)

type MallOrder struct {
	ID          int       `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'id'"`  // id
	UserID      int       `gorm:"index:idx_user_id;column:user_id;type:bigint;not null;default:0;comment:'id'"`        // id
	OrderID     string    `gorm:"index:idx_orderid;column:order_id;type:varchar(64);not null;default:'';comment:'id'"` // id
	Status      int       `gorm:"column:status;type:tinyint;not null;default:0;comment:'1 2 3 4 11 12'"`               // 1 2 3 4 11 12
	PayID       string    `gorm:"column:pay_id;type:varchar(100);not null;default:'';comment:'id'"`                    // id
	PayStatus   int       `gorm:"column:pay_status;type:tinyint;not null;default:0;comment:'0 1 2 3 4:'"`              // 0 1 2 3 4:
	PayType     int       `gorm:"column:pay_type;type:tinyint unsigned;not null;default:0;comment:'1,2,3,4,5'"`        // 1,2,3,4,5
	Source      int       `gorm:"column:source;type:tinyint;not null;default:0;comment:'1 2PC'"`                       // 1 2PC
	Serial      string    `gorm:"unique;column:serial;type:varchar(64);not null;default:''"`
	TotalAmount int       `gorm:"column:total_amount;type:int;not null;default:0;comment:','"`       // ,
	GoodsNum    int       `gorm:"column:goods_num;type:int unsigned;not null;default:1;comment:'1'"` // 1
	PayTime     time.Time `gorm:"column:pay_time;type:datetime;not null;default:1000-10-10 10:00:00"`
	CancelTime  time.Time `gorm:"column:cancel_time;type:datetime;not null;default:1000-10-10 10:00:00"`
	OrderExpand string    `gorm:"column:order_expand;type:varchar(512);not null;default:''"`
	CreatedAt   time.Time `gorm:"index:idx_created_at;column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	DeleteAt    time.Time `gorm:"column:deleted_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

const TABLE_MALL_ORDER = "mall_order"

type OrderDbDao struct {
	Db *gorm.DB
}

func NewOrderDbDao() OrderDbDao {
	r := OrderDbDao{
		Db: GetDbInstance(""),
	}
	return r
}

func (dao OrderDbDao) WithDBInstance(db *gorm.DB) OrderDbDao {
	dao.Db = db
	return dao
}

func (dao OrderDbDao) CreateOrder(ctx context.Context, order entity.OrderInfo) (orderId int, err error) {
	mallOrder := entityToDbOrder(order)
	db := dao.Db.Table(TABLE_MALL_ORDER).WithContext(ctx).Create(&mallOrder)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Create err:%v", err)
		return 0, db.Error
	}
	return mallOrder.ID, nil
}

func (dao OrderDbDao) GetOrderInfo(ctx context.Context, orderId string) (order entity.OrderInfo, err error) {
	var mallOrder MallOrder
	db := dao.Db.Table(TABLE_MALL_ORDER).WithContext(ctx).Where("order_id = ?", orderId).First(&mallOrder)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.First err:%v", err)
		return order, db.Error
	}
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return order, nil
	}
	order = dbOrderToEntity(mallOrder)
	return
}

func (dao OrderDbDao) FindOrderList(ctx context.Context, userId int) (orderList []entity.OrderInfo, err error) {
	orderList = make([]entity.OrderInfo, 0)
	var mallOrder []MallOrder
	db := dao.Db.Table(TABLE_MALL_ORDER).WithContext(ctx).Where("user_id = ?", userId).Find(&mallOrder)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", err)
		return orderList, db.Error
	}
	for _, v := range mallOrder {
		orderList = append(orderList, dbOrderToEntity(v))
	}
	return
}

func entityToDbOrder(en entity.OrderInfo) MallOrder {
	return MallOrder{
		OrderID:     en.OrderId,
		UserID:      en.UserId,
		GoodsNum:    en.GoodsNum,
		TotalAmount: en.TotalAmount,
		Serial:      en.Serial,
		Status:      en.Status,
		PayID:       en.PayId,
		PayStatus:   en.PayStatus,
		PayType:     en.PayType,
		PayTime:     en.PayTime,
	}
}

func dbOrderToEntity(order MallOrder) entity.OrderInfo {
	return entity.OrderInfo{
		Id:          int(order.ID),
		OrderId:     order.OrderID,
		UserId:      order.UserID,
		GoodsNum:    order.GoodsNum,
		TotalAmount: order.TotalAmount,
		Serial:      order.Serial,
		Status:      order.Status,
		PayId:       order.PayID,
		PayStatus:   order.PayStatus,
		PayType:     order.PayType,
		PayTime:     order.PayTime,
	}
}
