package repo

import (
	"context"
	"mall/internal/dao/db"
	"mall/internal/entity"
)

type UserRepositoryImpl struct {
	userDao db.UserDbDao
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{
		userDao: db.NewUserDbDao(),
	}
}

func (repo UserRepositoryImpl) CreateUser(ctx context.Context, user entity.User) (rowId int, err error) {
	return repo.userDao.CreateUser(ctx, user)
}
func (repo UserRepositoryImpl) GetUserByAccount(ctx context.Context, account, pwd string) (user entity.User, err error) {
	return repo.userDao.GetUserByAccount(ctx, account, pwd)
}
func (repo UserRepositoryImpl) FindUserById(ctx context.Context, userId int) (user entity.User, err error) {
	return repo.userDao.FindUserById(ctx, userId)
}
