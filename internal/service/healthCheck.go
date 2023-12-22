package service

import (
	"context"
	"mall/internal/logger"
	"math/rand"
	"time"
)

func HealthCheck(ctx context.Context) (resp interface{}, err error) {
	r := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(r))
	logger.WithContext(ctx).Infof("service.HealthCheck ret:%v", r)
	return r, nil
}

func HealthCheckV1() (resp interface{}, err error) {
	r := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(r))
	logger.WithGoID().Infof("service.HealthCheck ret:%v", r)
	return r, nil
}
