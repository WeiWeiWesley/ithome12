package main

import (
	"fmt"
	tarce "ithome12/basicGo/tarceInit"
)

const mainConst = "mainConst"
var mainVar string

//會在 package init()完成後執行
func init() {
	mainVar = "mainVar"
	fmt.Println("mainConst", mainConst)
	fmt.Println("mainVar", mainVar)
}

func main() {
	fmt.Println("main func started")
	//不管呼叫幾次 trace package，trace init()都只會被執行一次
	tarce.Print("hello")
	tarce.Print("hello2")
}
