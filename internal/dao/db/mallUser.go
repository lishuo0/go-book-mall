package db

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"mall/internal/entity"
	"mall/internal/logger"
	"time"
)

type MallUser struct {
	ID        int       `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null;comment:'ID'"` // ID
	NickName  string    `gorm:"column:nick_name;type:varchar(64);not null;default:''"`
	Account   string    `gorm:"index:idx_account;column:account;type:varchar(32);not null;default:''"`
	Password  string    `gorm:"column:password;type:varchar(64);not null;default:''"`
	Icon      string    `gorm:"column:icon;type:varchar(256);not null;default:''"`
	Gender    int       `gorm:"column:gender;type:tinyint unsigned;not null;default:0;comment:'12'"`  // 12
	Status    int       `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:'123'"` // 123
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:timestamp;default:null"`
}

const TABLE_MALL_USER = "mall_user"

type UserDbDao struct {
	Db *gorm.DB
}

func NewUserDbDao() UserDbDao {
	r := UserDbDao{
		Db: GetDbInstance(""),
	}
	return r
}

func (dao UserDbDao) WithDBInstance(db *gorm.DB) UserDbDao {
	dao.Db = db
	return dao
}

func (dao UserDbDao) CreateUser(ctx context.Context, user entity.User) (rowId int, err error) {
	mallUser := entityToDbUser(user)
	db := dao.Db.Table(TABLE_MALL_USER).WithContext(ctx).Create(&mallUser)
	if db.Error != nil {
		logger.WithContext(ctx).Errorf("db create err:%v", err)
		return 0, db.Error
	}

	return cast.ToInt(mallUser.ID), nil
}

func (dao UserDbDao) GetUserByAccount(ctx context.Context, account, pwd string) (user entity.User, err error) {
	var mallUser MallUser
	db := dao.Db.Table(TABLE_MALL_USER).WithContext(ctx).Where("account = ?", account).Where("password = ?", pwd).Find(&mallUser)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db First err:%v", err)
		return user, db.Error
	}

	return dbUserToEntity(mallUser), nil
}

func (dao UserDbDao) FindUserById(ctx context.Context, userId int) (user entity.User, err error) {
	var mallUser MallUser
	mallUser.ID = userId
	db := dao.Db.Table(TABLE_MALL_USER).WithContext(ctx).First(&mallUser)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db first err:%v", err)
		return user, db.Error
	}

	return dbUserToEntity(mallUser), nil
}

// 业务实体：代码内部使用的；数据库结构定义的一般不一样。
func dbUserToEntity(db MallUser) entity.User {
	return entity.User{
		Id:        cast.ToInt(db.ID),
		NickName:  db.NickName,
		Account:   db.Account,
		Password:  db.Password,
		Icon:      db.Icon,
		Gender:    db.Gender,
		Status:    db.Status,
		CreateAt:  db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
	}
}

func entityToDbUser(en entity.User) MallUser {
	mallUser := MallUser{
		ID:       en.Id,
		NickName: en.NickName,
		Account:  en.Account,
		Password: en.Password,
		Icon:     en.Icon,
		Gender:   en.Gender,
		Status:   en.Status,
	}
	return mallUser
}
