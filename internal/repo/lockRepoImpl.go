package repo

import (
	"context"
	"mall/internal/dao/redis"
	"time"
)

type DistributeLockRepositoryImpl struct {
	lockDao redis.DistributeLockDao
}

func NewDistributeLockRepository() DistributeLockRepository {
	return &DistributeLockRepositoryImpl{
		lockDao: redis.NewDistributeLockDao(),
	}
}

func (repo *DistributeLockRepositoryImpl) Lock(ctx context.Context, key string, expire time.Duration) bool {
	return repo.lockDao.Lock(ctx, key, expire)
}
func (repo *DistributeLockRepositoryImpl) UnLock(ctx context.Context, key string) bool {
	return repo.lockDao.UnLock(ctx, key)
}
