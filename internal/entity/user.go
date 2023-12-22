package entity

import "time"

// https://pkg.go.dev/github.com/go-playground/validator/v10
type SetUserInfoReq struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName" binding:"min=3,max=64"`
	Gender   int    `json:"gender" binding:"oneof=1 2"`
	Age      int    `json:"age" binding:"min=1,max=100"`
}

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

type UserListReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserListResp struct {
	Users []User `json:"users"`
}

type User struct {
	Id        int       `json:"id"`
	NickName  string    `json:"nick_name"`
	Account   string    `json:"account"`
	Password  string    `json:"-"` // 数据库存储，需要加密
	Icon      string    `json:"icon"`
	Gender    int       `json:"gender"`
	Status    int       `json:"status"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
