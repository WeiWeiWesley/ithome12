package main

import "fmt"

//Burger 就是個漢堡
type Burger struct {
	beef    bool
	cheese  bool
	pickles bool
}

var (
	addBeff    = make(chan Burger, 10) //超過10筆後阻塞
	addCheese  = make(chan Burger, 10) //超過10筆後阻塞
	addPickles = make(chan Burger, 10) //超過10筆後阻塞
	done       = make(chan []Burger)
)

func main() {
	go AddBeff()    //加牛肉監聽開始
	go AddCheese()  //加起司監聽開始
	go AddPickles() //加酸黃瓜監聽開始

	for {
		select {
		//達成後停止
		case count := <-done:
			for i := range count {
				if !count[i].beef || !count[i].cheese || !count[i].pickles {
					fmt.Println("oops!", count)
					return
				}
			}
			fmt.Println(len(count), "success!")
			return
		//開始製造漢堡
		case addBeff <- Burger{}:
		}
	}
}

//AddBeff 加牛肉
func AddBeff() {
	for {
		//一次只能加一塊牛肉
		select {
		case burger := <-addBeff:
			burger.beef = true
			addCheese <- burger
		}
	}
}

//AddCheese 加起司
func AddCheese() {
	for {
		//一次只能加一塊起司
		select {
		case burger := <-addCheese:
			burger.cheese = true
			addPickles <- burger
		}
	}
}

//AddPickles 加酸黃瓜
func AddPickles() {
	var count []Burger
	for {
		//一次只能加一份酸黃瓜
		select {
		case burger := <-addPickles:
			burger.pickles = true
			//漢堡完成
			count = append(count, burger)

			//達成目標
			if len(count) == 1000 {
				done <- count
			}
		}
	}
}
