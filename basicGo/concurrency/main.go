package main

import (
	"fmt"
	"time"
)

func main() {
	serialize()
	fmt.Println("-----")
	concurrency()
	time.Sleep(time.Second)
}

func serialize() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) //為方便觀察睡一下
	}
}

func concurrency() {
	for i := 0; i < 5; i++ {
		//golang 併發的關鍵保留字 "go"
		go func(i int) {
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond) //由於是同時進行，0~4將同時 sleep
		}(i)
	}
}
