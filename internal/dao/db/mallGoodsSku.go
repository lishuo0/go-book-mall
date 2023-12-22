package db

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mall/internal/entity"
	"mall/internal/logger"
	"strings"
	"time"
)

type MallGoodsSku struct {
	ID            int       `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null;comment:'id'"`      // id
	GoodsID       int       `gorm:"index:idx_goods_id;column:goods_id;type:int unsigned;not null;default:0;comment:'id'"` // id
	AttributeIDs  string    `gorm:"column:attribute_ids;type:varchar(255);not null;default:'';comment:'skuid, '"`         // skuid,
	SpendPrice    int       `gorm:"column:spend_price;type:int unsigned;not null;default:0"`
	Price         int       `gorm:"column:price;type:int unsigned;not null;default:0"`
	DiscountPrice int       `gorm:"column:discount_price;type:int unsigned;not null;default:0"`
	LeftStore     int       `gorm:"column:left_store;type:int unsigned;not null;default:0"`
	AllStore      int       `gorm:"column:all_store;type:int unsigned;not null;default:0"`
	Status        int       `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:'1: ; 2: '"` // 1: ; 2:
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

const TABLE_MALL_GOODS_SKU = "mall_goods_sku"

type GoodsSkuDbDao struct {
	Db *gorm.DB
}

func NewGoodsSkuDbDao() GoodsSkuDbDao {
	r := GoodsSkuDbDao{
		Db: GetDbInstance(""),
	}
	return r
}

func (dao GoodsSkuDbDao) WithDBInstance(db *gorm.DB) GoodsSkuDbDao {
	dao.Db = db
	return dao
}

func (dao GoodsSkuDbDao) CreateGoodsSku(ctx context.Context, goodsSku []entity.GoodsSkuInfo) (affectd int, err error) {
	var mallSkus []MallGoodsSku
	for _, v := range goodsSku {
		mallSkus = append(mallSkus, entityToDbSku(v))
	}
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).CreateInBatches(mallSkus, 2)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Create err:%v", db.Error)
		return 0, db.Error
	}

	return int(db.RowsAffected), nil
}

func (dao GoodsSkuDbDao) FindGoodsSkuByGoodId(ctx context.Context, goodsId int) (goodSkuList []entity.GoodsSkuInfo, err error) {
	goodSkuList = make([]entity.GoodsSkuInfo, 0)
	var mallSkus []MallGoodsSku
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).Where("goods_id = ?", goodsId).Find(&mallSkus)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", db.Error)
		return goodSkuList, db.Error
	}
	for _, v := range mallSkus {
		goodSkuList = append(goodSkuList, dbSkuToEntity(v))
	}

	return
}

func (dao GoodsSkuDbDao) SelectGoodsBySkuId(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {
	var mallSku MallGoodsSku
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).Where("id = ?", skuId).First(&mallSku)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", db.Error)
		return goodSku, db.Error
	}
	goodSku = dbSkuToEntity(mallSku)

	return
}

func (dao GoodsSkuDbDao) SelectGoodsBySkuIdForUpdate(ctx context.Context, skuId int) (goodSku entity.GoodsSkuInfo, err error) {
	var mallSku MallGoodsSku
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", skuId).First(&mallSku)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", db.Error)
		return goodSku, db.Error
	}
	goodSku = dbSkuToEntity(mallSku)

	return
}

func (dao GoodsSkuDbDao) FindGoodsSkuBySkuIds(ctx context.Context, skuIds []int) (goodSkuList []entity.GoodsSkuInfo, err error) {
	goodSkuList = make([]entity.GoodsSkuInfo, 0)
	var mallSkus []MallGoodsSku
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).Where("id in ?", skuIds).Find(&mallSkus)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", db.Error)
		return goodSkuList, db.Error
	}
	for _, v := range mallSkus {
		goodSkuList = append(goodSkuList, dbSkuToEntity(v))
	}

	return
}

func (dao GoodsSkuDbDao) UpdateGoodsSkuStore(ctx context.Context, sku entity.GoodsSkuInfo) (affectd int, err error) {
	db := dao.Db.Table(TABLE_MALL_GOODS_SKU).WithContext(ctx).Where("id = ?", sku.Id).Updates(MallGoodsSku{LeftStore: sku.Leftstore})
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Updates err:%v", db.Error)
		return 0, db.Error
	}

	return cast.ToInt(db.RowsAffected), nil
}

func entityToDbSku(info entity.GoodsSkuInfo) MallGoodsSku {
	attIds := make([]string, 0)
	for _, v := range info.AttIds {
		attIds = append(attIds, cast.ToString(v))
	}
	return MallGoodsSku{
		GoodsID:       info.GoodsId,
		AttributeIDs:  strings.Join(attIds, ","),
		SpendPrice:    info.SpendPrice,
		Price:         info.Price,
		DiscountPrice: info.DiscountPrice,
		AllStore:      info.Allstore,
		LeftStore:     info.Leftstore,
	}
}

func dbSkuToEntity(info MallGoodsSku) entity.GoodsSkuInfo {
	attIds := make([]int, 0)
	attArr := strings.Split(info.AttributeIDs, ",")
	for _, v := range attArr {
		attIds = append(attIds, cast.ToInt(v))
	}

	return entity.GoodsSkuInfo{
		Id:            info.ID,
		GoodsId:       info.GoodsID,
		AttIds:        attIds,
		SpendPrice:    info.SpendPrice,
		Price:         info.Price,
		DiscountPrice: info.DiscountPrice,
		Allstore:      info.AllStore,
		Leftstore:     info.LeftStore,
	}
}
