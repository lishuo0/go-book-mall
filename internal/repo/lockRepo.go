package repo

import (
	"context"
	"time"
)

type DistributeLockRepository interface {
	Lock(ctx context.Context, key string, expire time.Duration) bool
	UnLock(ctx context.Context, key string) bool
}
