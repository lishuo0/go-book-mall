package service

import (
	"context"
	"mall/api/httputils"
	"mall/internal/constant"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/repo"
	"math/rand"
	"strconv"
)

func APILogin(ctx context.Context, req entity.LoginReq) (resp entity.LoginResp, err error) {

	lockRepo := repo.NewDistributeLockRepository()
	suc := lockRepo.Lock(ctx, req.Account, constant.RedisDistributeLockExpire)
	if !suc {
		return resp, httputils.UserLoginError
	}
	userRepo := repo.NewUserRepository()
	user, err := userRepo.GetUserByAccount(ctx, req.Account, req.Password)
	if err != nil {
		logger.WithContext(ctx).Errorf("userRepo.GetUserByAccount err:%v", err)
		return resp, err
	}
	// 没查到用户说明第一次登录，新增用户
	if user.Id == 0 {
		user.Account = req.Account
		user.Password = req.Password
		user.NickName = "匿名用户" + strconv.Itoa(rand.Intn(10000))
		user.Id, err = userRepo.CreateUser(ctx, user)
	}

	lockRepo.UnLock(ctx, req.Account)

	token, _ := CreateAPIToken(0, constant.LoginTokenExpireDefault)

	resp.UserId = user.Id
	resp.Token = token
	return resp, nil
}

func UserList(ctx context.Context, req entity.LoginReq) (resp entity.LoginResp, err error) {

	userRepo := repo.NewUserRepository()
	user, err := userRepo.GetUserByAccount(ctx, req.Account, req.Password)
	if err != nil {
		logger.WithContext(ctx).Errorf("userRepo.GetUserByAccount err:%v", err)
		return resp, err
	}
	// 没查到用户说明第一次登录，新增用户
	if user.Id == 0 {
		user.Account = req.Account
		user.Password = req.Password
		user.NickName = "匿名用户" + strconv.Itoa(rand.Intn(10000))
		user.Id, err = userRepo.CreateUser(ctx, user)
	}

	token, _ := CreateAPIToken(0, constant.LoginTokenExpireDefault)

	resp.Token = token
	return resp, nil
}

func SetUserInfo(ctx context.Context, userInfo entity.SetUserInfoReq) error {
	return nil
}
