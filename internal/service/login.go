package service

import (
	"github.com/dgrijalva/jwt-go"
	"mall/internal/constant"
	"time"
)

func CreateAPIToken(userId int, failTokenMin time.Duration) (token string, err error) {
	// jwt 生成token
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(failTokenMin).Unix(),
	})

	// 随机密钥
	atoken, err := at.SignedString([]byte(constant.APILoginSecretKey))
	if err != nil {
		return "", err
	}
	return atoken, nil
}

func ParseAPIToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.APILoginSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
