package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mall/internal/entity"
	"mall/internal/logger"
	"time"
)

type MallGoods struct {
	ID          int       `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null;comment:'ID'"` // ID
	Name        string    `gorm:"column:name;type:varchar(100);not null;default:''"`
	Description string    `gorm:"column:description;type:varchar(255);not null;default:''"`
	Tags        string    `gorm:"column:tags;type:varchar(255);not null;default:''"`
	Detail      string    `gorm:"column:detail;type:text;not null"`
	CategoryID  int       `gorm:"column:category_id;type:int;not null;default:0"`
	SmallImage  string    `gorm:"column:small_image;type:varchar(255);not null;default:''"`
	DetailImage string    `gorm:"column:detail_image;type:varchar(255);not null;default:''"`
	Price       int       `gorm:"column:price;type:int;not null;default:0"`
	Status      int       `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:'1: ; 2: '"` // 1: ; 2:
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

const TABLE_MALL_GOODS = "mall_goods"

type GoodsDbDao struct {
	Db *gorm.DB
}

func NewGoodsDbDao() GoodsDbDao {
	r := GoodsDbDao{
		Db: GetDbInstance(""),
	}
	return r
}

func (dao GoodsDbDao) WithDBInstance(db *gorm.DB) GoodsDbDao {
	dao.Db = db
	return dao
}

func (dao GoodsDbDao) CreateGoods(ctx context.Context, goods entity.GoodsInfo) (goodId int, err error) {
	// 业务实体，与底层数据库表不一致，转换
	var mallGoods MallGoods
	mallGoods = entityToDbGoods(goods)
	db := dao.Db.Table(TABLE_MALL_GOODS).WithContext(ctx).Create(&mallGoods)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Create err:%v", db.Error)
		return 0, db.Error
	}

	return mallGoods.ID, nil
}

func (dao GoodsDbDao) UpdateGoods(ctx context.Context, goods entity.GoodsInfo) (goodId int, err error) {
	var mallGoods MallGoods
	mallGoods = entityToDbGoods(goods)
	db := dao.Db.Table(TABLE_MALL_GOODS).WithContext(ctx).Where("id = ?", 6).Updates(&mallGoods)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Updates err:%v", db.Error)
		return 0, db.Error
	}

	return int(db.RowsAffected), nil
}

func (dao GoodsDbDao) DeleteGoods(ctx context.Context, goods entity.GoodsInfo) (goodId int, err error) {
	var mallGoods = MallGoods{
		ID: 28,
	}
	//mallGoods = entityToDbGoods(goods)
	db := dao.Db.Table(TABLE_MALL_GOODS).WithContext(ctx).Delete(&mallGoods)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db.Delete err:%v", db.Error)
		return 0, db.Error
	}

	return int(db.RowsAffected), nil
}

func (dao GoodsDbDao) FindGoodsListByCategoryId(ctx context.Context, cateId int) (goodsList []entity.GoodsInfo, err error) {
	goodsList = make([]entity.GoodsInfo, 0)
	var mallGoods []MallGoods
	db := dao.Db.Table(TABLE_MALL_GOODS).WithContext(ctx).Where("category_id = ?", cateId).
		Order("updated_at desc").Limit(5).Offset(10).Find(&mallGoods)

	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.Find err:%v", db.Error)
		return goodsList, db.Error
	}
	for _, v := range mallGoods {
		goodsList = append(goodsList, dbGoodsToEntity(v))
	}

	return
}

func (dao GoodsDbDao) GetGoodsInfoById(ctx context.Context, goodsId int) (goods entity.GoodsInfo, err error) {
	var mallgoods MallGoods
	mallgoods.ID = goodsId
	db := dao.Db.Table(TABLE_MALL_GOODS).WithContext(ctx).First(&mallgoods)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db.First err:%v", db.Error)
		return goods, db.Error
	}
	goods = dbGoodsToEntity(mallgoods)
	return
}

func entityToDbGoods(info entity.GoodsInfo) MallGoods {
	return MallGoods{
		ID:          info.Id,
		CategoryID:  info.CategoryId,
		Name:        info.Name,
		Description: info.Description,
		SmallImage:  info.SmallImage,
		DetailImage: info.DetailImage,
		Status:      info.Status,
	}
}

func dbGoodsToEntity(info MallGoods) entity.GoodsInfo {
	return entity.GoodsInfo{
		Id:          info.ID,
		CategoryId:  info.CategoryID,
		Name:        info.Name,
		Description: info.Description,
		SmallImage:  info.SmallImage,
		DetailImage: info.DetailImage,
		Status:      info.Status,
	}
}
