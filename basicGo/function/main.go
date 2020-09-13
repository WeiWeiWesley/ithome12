package main

import (
	"fmt"
	"strconv"
)

func main() {
	//單一參數，無回傳
	voidFunc("Wesley")

	//多參數，多回傳
	sum, err := intSum("1", "1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(sum)

	//匿名函式，常配合併發使用
	func(f float64) {
		fmt.Println("π", f)
	}(3.1415926)
}

//單一參數，無回傳
func voidFunc(name string) {
	fmt.Println(name)
}

//多參數，多回傳
func intSum(numA string, numB string) (int, error) {
	a, err := strconv.Atoi(numA)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	b, err := strconv.Atoi(numB)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return a + b, nil
}
