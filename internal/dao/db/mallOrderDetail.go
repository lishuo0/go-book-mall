package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mall/internal/entity"
	"mall/internal/logger"
	"time"
)

type MallOrderDetail struct {
	ID          int       `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'id'"`                                // id
	OrderID     string    `gorm:"uniqueIndex:idx_order_sku_id;index:idx_order_id;column:order_id;type:varchar(64);not null;default:'';comment:'id'"` // id
	SkuID       int       `gorm:"uniqueIndex:idx_order_sku_id;index:idx_sku_id;column:sku_id;type:int;not null;default:0;comment:'sku_id'"`          // sku_id
	Price       int       `gorm:"column:price;type:int unsigned;not null;default:1"`
	Num         int       `gorm:"column:num;type:int unsigned;not null;default:1;comment:'1'"` // 1
	CreatedAt   time.Time `gorm:"index:idx_created_at;column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	OrderExpand string    `gorm:"column:order_expand;type:varchar(512);not null;default:''"`
	Status      int       `gorm:"column:status;type:tinyint;not null;default:1;comment:'1 2'"` // 1 2
}

const TABLE_MALL_ORDER_DETAIL = "mall_order_detail"

type OrderDetailDbDao struct {
	Db *gorm.DB
}

func NewOrderDetailDbDao() OrderDetailDbDao {
	r := OrderDetailDbDao{
		Db: GetDbInstance(""),
	}
	return r
}

func (dao OrderDetailDbDao) WithDBInstance(db *gorm.DB) OrderDetailDbDao {
	dao.Db = db
	return dao
}

func (dao OrderDetailDbDao) CreateOrderDetails(ctx context.Context, orderDetails []entity.OrderDetailInfo) (err error) {
	var mallDetails []MallOrderDetail
	for _, v := range orderDetails {
		mallDetails = append(mallDetails, entityToDbOrderDetail(v))
	}
	db := dao.Db.Table(TABLE_MALL_ORDER_DETAIL).WithContext(ctx).Create(&mallDetails)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Create err:%v", err)
		return db.Error
	}

	return
}

func (dao OrderDetailDbDao) FindOrderDetailList(ctx context.Context, orderId string) (detailList []entity.OrderDetailInfo, err error) {
	detailList = make([]entity.OrderDetailInfo, 0)
	var mallDetails []MallOrderDetail
	db := dao.Db.Table(TABLE_MALL_ORDER_DETAIL).WithContext(ctx).Where("order_id = ?", orderId).Find(&mallDetails)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", err)
		return detailList, db.Error
	}
	for _, v := range mallDetails {
		detailList = append(detailList, dbOrderDetailToEntity(v))
	}
	return
}

func entityToDbOrderDetail(en entity.OrderDetailInfo) MallOrderDetail {
	return MallOrderDetail{
		OrderID: en.OrderId,
		SkuID:   en.SkuId,
		Price:   en.Price,
		Num:     en.Num,
	}
}

func dbOrderDetailToEntity(order MallOrderDetail) entity.OrderDetailInfo {
	return entity.OrderDetailInfo{
		Id:      order.ID,
		OrderId: order.OrderID,
		SkuId:   order.SkuID,
		Price:   order.Price,
		Num:     order.Num,
		Status:  order.Status,
	}
}
