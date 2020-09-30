package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println(time.Now().String())
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(10) == 1 {
			panic("Now panic and restart")
		}
	}
}
