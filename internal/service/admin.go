package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"mall/internal/constant"
	"mall/internal/entity"
	"time"
)

func AdminLogin(ctx context.Context, req entity.AdminLoginReq) (resp entity.AdminLoginResp, err error) {

	token, _ := CreateAPIToken(0, constant.LoginTokenExpireDefault)

	resp.Token = token
	return resp, nil
}

func CreateAdminToken(adminId int, failTokenMin time.Duration) (token string, err error) {
	// jwt 生成token
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": adminId,
		"exp":      time.Now().Add(failTokenMin).Unix(),
	})

	atoken, err := at.SignedString([]byte(constant.LoginAdminSecretKey))
	if err != nil {
		return "", err
	}
	return atoken, nil
}

func ParseAdminToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.LoginAdminSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
