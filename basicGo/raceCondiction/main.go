package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// exampleRace()
	// exampleChannel()
	exampleMutex()
}

func exampleMutex() {
	lock := sync.Mutex{}
	count := 0

	for i := 0; i < 10; i++ {
		go func() {
			//在對相同記憶體位置讀寫時，上鎖
			lock.Lock()
			count++
			//在對相同記憶體位置讀寫完成時，解鎖
			lock.Unlock()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count)
}

func exampleChannel() {
	count := 0

	//宣告一個 no buffer channel
	//利用其阻塞特性達成等待
	blockCh := make(chan bool)

	//waitting channel
	go func() {
		//blockCh 收到值才會++
		for {
			select {
			case <-blockCh:
				count++
			}
		}
	}()

	//concurrency
	for i := 0; i < 10; i++ {
		go func() {
			blockCh <- true
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count)
}

func exampleRace() {
	count := 0
	for i := 0; i < 10; i++ {
		go func() {
			count++
		}()
	}
	fmt.Println(count)
}
