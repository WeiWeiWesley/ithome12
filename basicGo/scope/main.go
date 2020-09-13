package main

import (
	"fmt"
	"ithome12/basicGo/scpackage"
)

var global = "a global var"

func main() {
	//scope of package
	client := scpackage.Create("127.0.0.1", "localhost")
	fmt.Printf("%+v\n", client)

	//scope of block
	myFuncA()
	myFuncB()
	fmt.Println(global)

	new := "new"
	for i := 0; i < 1; i++ {
		//new 被宣告於迴圈{}外，符合作用域可使用
		//i 屬於此迴圈內宣告變數，故僅能作用於迴圈內
		//newInFor 同 i，屬於此迴圈內宣告變數，故僅能作用於迴圈內
		newInFor := "hi"
		fmt.Println(i, new, newInFor)
	}

	//每個block{}都分別為新的作用域
	//if 判斷式內的 "new" 變數與 先前的 "new" 變數，為不同實體，並有著不同記憶體位址
	if len(new) > 0 {
		new := "xxx"
		fmt.Println(new, &new)
	}
	fmt.Println(new, &new)
}

func myFuncA() {
	fmt.Println(global)
	global = "changed by myFuncA"
}

func myFuncB() {
	fmt.Println(global)
	global = "changed by myFuncB"
}
