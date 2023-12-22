package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"sync"
	"time"
)

var cache *bigcache.BigCache

var once sync.Once

func initBigCache() {
	cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(time.Minute*10))
}

// 单例初始化
func GetBigCacheInstance() *bigcache.BigCache {
	if cache != nil {
		return cache
	}

	once.Do(func() {
		initBigCache()
	})

	return cache
}
