package service

import (
	"context"
	"encoding/json"
	"mall/internal/entity"
)

type OrderCreateJob struct {
	Req    entity.CreateOrderReq
	DoneC  chan struct{}
	ErrorC chan error
}

func NewOrderCreateJob(req entity.CreateOrderReq) *OrderCreateJob {
	return &OrderCreateJob{
		Req:    req,
		DoneC:  make(chan struct{}, 0),
		ErrorC: make(chan error, 1),
	}
}

func (job *OrderCreateJob) Execute(ctx context.Context) (err error) {
	_, err = CreateOrder(ctx, job.Req)
	return
}

func (job *OrderCreateJob) String() (str string) {
	data, _ := json.Marshal(job)
	return string(data)
}

// close 实现信号的传递
func (job *OrderCreateJob) Done() {
	close(job.DoneC)
}

// 写到管道，非阻塞写
func (job *OrderCreateJob) Error(err error) {
	select {
	case job.ErrorC <- err:
	default:
	}
}
