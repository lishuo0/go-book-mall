package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var count int32
	var lock sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			count++
			lock.Unlock()
		}
	})
}

func BenchmarkCas(b *testing.B) {
	var count int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for !atomic.CompareAndSwapInt32(&count, count, count+1) {
			}
		}
	})
}
