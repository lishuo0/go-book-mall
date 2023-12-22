package service

import (
	"context"
	"errors"
	"github.com/hashicorp/go-uuid"
	"mall/internal/entity"
	"mall/internal/logger"
	"sync"
	"time"
)

var (
	once  = sync.Once{}
	queue *AysnQueue
)

// 异步队列：订单，难以满足，可靠性不行
type AysnQueue struct {
	// 管道
	Queue chan entity.AsynQueueJob

	// 令牌桶，k可以存储很多令牌
	Bucket chan struct{}
}

// 消费：控制速率，慢一点
func NewAsynQueue(limit, bucket int) *AysnQueue {
	if queue != nil {
		return queue
	}

	once.Do(func() {
		queue = &AysnQueue{
			// 管道大小 2048：管道满了怎么办？大小根据业务情况初始化
			Queue:  make(chan entity.AsynQueueJob, 2048),
			Bucket: make(chan struct{}, bucket),
		}

		// 限流：定时产生令牌的方式；Bucket 存储一些令牌，目前没消费请求，缓存一些令牌；待会如果有突发请求，可以处理
		// 队列消费，如果有限流：加了定时器
		ticker := time.NewTicker(time.Second / time.Duration(limit))
		go func() {
			for {
				<-ticker.C //时间触发时，管道可读
				select {
				case queue.Bucket <- struct{}{}:
				default:
				}
			}
		}()

		// 消费协程数，根据业务情况
		for i := 0; i < 10; i++ {
			go func() {
				for {
					job := <-queue.Queue

					//要处理请求，获取令牌才能处理
					<-queue.Bucket

					traceId, _ := uuid.GenerateUUID()
					ctx := context.WithValue(context.Background(), "traceId", traceId)
					logger.WithContext(ctx).Errorf("consumer asyn_job:%s", job.String())
					err := job.Execute(ctx)
					if err != nil {
						logger.WithContext(ctx).Errorf("job.Execute error:%v", err)
						// 异步通知错误
						job.Error(err)
						continue
					}

					// 异步通知 完成
					job.Done()
				}
			}()
		}
	})

	return queue
}

func (q *AysnQueue) PushDelayJob(ctx context.Context, job entity.AsynQueueJob, delay time.Duration) {
	time.AfterFunc(delay, func() {
		q.Queue <- job
	})
}

// 非阻塞写，不会阻塞处理协程
func (q *AysnQueue) PushJob(ctx context.Context, job entity.AsynQueueJob) error {
	select {
	case q.Queue <- job:
		return nil
	default:
		logger.WithContext(ctx).Errorf("push job to queue block:%v", job)
		return errors.New("push job to queue block")
	}
}
