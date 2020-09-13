package main

import "fmt"

func main() {
	//指定範圍迴圈
	fmt.Println("count 0~9")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	//foreach
	fmt.Println("foreach")
	memberList := []string{"Wesley", "Ken", "Eric"} //a string slice
	for key, value := range memberList {
		fmt.Println(key, value)
	}
	fmt.Println()

	//break loop
	fmt.Println("break loop when i over 3")
	for i := 0; i < 10; i++ {
		if i > 3 {
			fmt.Println("break at", i)
			break
		}
		fmt.Println(i)
	}
	fmt.Println()

	//continue
	fmt.Println("use continue to skip 5")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //will skip code after "continue" and keep going next round
		}
		fmt.Println(i)
	}
	fmt.Println()
}
