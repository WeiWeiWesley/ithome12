package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	var count int32
	for i := 0; i < b.N; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
		}()
	}
}

func BenchmarkMutex(b *testing.B) {
	lock := sync.Mutex{}
	var count int32

	for i := 0; i < b.N; i++ {
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
		}()
	}
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int32)
	var count int32

	go func() {
		for {
			select {
			case add := <-ch:
				count += add
			}
		}

	}()

	for i := 0; i < b.N; i++ {
		go func() {
			ch <- 1
		}()
	}
}
