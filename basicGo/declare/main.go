package main

//套件引用
import (
	"fmt"
	"time"
)

func main() {
	//基本宣告，給予該型態預設值
	var number int
	var str string

	//多變數宣告
	var (
		numFloat float32
		numInt64 int64
	)

	//都會得到該型態預設值
	fmt.Println("number", number)
	fmt.Println("str", str)
	fmt.Println("numFloat", numFloat)
	fmt.Println("numInt64", numInt64)

	//短宣告，可同時賦值，且型態將由賦值時自動定義
	name := "Wesley"
	age := 32
	createTime := time.Now()

	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("createTime", createTime)

	//部分變數在賦值前需 make 出來
	var strMap map[string]string
	strMap = make(map[string]string) //不信你可以註解掉這行
	strMap["ip"] = "127.0.0.1"
	strMap["host"] = "localhost"
	fmt.Println("strMap", strMap)
}
