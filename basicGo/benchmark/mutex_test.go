package benchmark_test

import (
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	lock := sync.Mutex{}
	count := 0

	for i := 0; i < b.N; i++ {
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
		}()
	}
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int)
	count := 0

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
