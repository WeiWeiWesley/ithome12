package tarce

import "fmt"

const packageConst = "packageConst"

var packageVar string

func init() {
	packageVar = "packageVar"
	fmt.Println("packageConst", packageConst)
	fmt.Println("packageVar", packageVar)
}

//Print something
func Print(s string) {
	fmt.Println(s)
}