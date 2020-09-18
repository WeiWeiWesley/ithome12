package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	exampleCount()
}

func exampleCount() {
	var count int64
	timeStop := time.NewTimer(3 * time.Second)

	//每100ms讀寫一次
	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				atomic.AddInt64(&count, 1)
				fmt.Println(count)
			}
		}
	}()

	//每100ms讀寫一次
	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C:
				atomic.AddInt64(&count, 1)
				fmt.Println(count)
			}
		}
	}()

	select {
	//預計執行 3s 60次
	case <-timeStop.C:
		time.Sleep(time.Millisecond) //防止最後一筆來不及寫入
		if count == 60 {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}
