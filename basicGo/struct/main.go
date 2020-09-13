package main

import "fmt"

//Client 定義 Client 型態
type Client struct {
	IP      string
	Host    string
	maxConn int
}

func main() {
	//回傳型態為 Client 的 struct
	client := newClient("127.0.0.1", "localhost")
	fmt.Printf("%+v\n", client)

	//可對已實體化的 client 賦值
	client.IP = "0.0.0.0"
	fmt.Printf("%+v\n", client)

	//可利用 pointer 的方式，pass by address
	setMaxConn(&client, 100)
	fmt.Printf("%+v\n", client)

	//直接用短宣告實體化
	client2 := Client{
		IP:      "192.0.0.1",
		Host:    "test.host",
		maxConn: 3,
	}
	fmt.Printf("%+v\n", client2)
}

//可用於傳遞
func newClient(ip, host string) Client {
	return Client{
		IP:      ip,
		Host:    host,
		maxConn: 10,
	}
}

//可利用 pointer pass by address 並異動
func setMaxConn(c *Client, num int) {
	c.maxConn = num
}
