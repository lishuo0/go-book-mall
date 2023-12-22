package test

import (
	"context"
	"mall/internal/core"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/service"
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	_ = core.InitConfig("/Users/lile/Documents/gocode/go-class/mall/configs/conf.yaml")
	_ = logger.InitLogger()

	login := entity.LoginReq{
		Account:  "test-user-3",
		Password: "123456",
	}
	var user1, user2 entity.LoginResp
	var err error

	go func() {
		user1, err = service.APILogin(context.Background(), login)
		if err != nil {
			t.Error(err)
		}

	}()
	go func() {
		user2, err = service.APILogin(context.Background(), login)
		if err != nil {
			t.Error(err)
		}

	}()
	time.Sleep(time.Second * 1)
	if user1.UserId != user2.UserId {
		t.Errorf("user1.UserId=%d not equal user2.UserId=%d", user1.UserId, user2.UserId)
	}

}
