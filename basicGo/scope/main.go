package main

import (
	"fmt"
	"ithome12/basicGo/scpackage"
)


func main() {
	client := scpackage.Create("127.0.0.1", "localhost")

	fmt.Printf("%+v\n", client)
}
