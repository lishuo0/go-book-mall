package entity

import "context"

// 对象：执行
type AsynQueueJob interface {
	Execute(ctx context.Context) error
	String() string
	Done()
	Error(err error)
}
