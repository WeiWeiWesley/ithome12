package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type adCache struct {
	lock  sync.RWMutex
	cache map[string]string
}

var (
	ad     string
	adList adCache
)

func init() {
	ad = "ad:2020-09"
	adList.lock = sync.RWMutex{}
	adList.cache = make(map[string]string)
}

func main() {
	// redisExample()
	localCache()
}

func localCache() {
	adList.lock.RLock()
	if res, ok := adList.cache[ad]; ok {
		fmt.Println("success", res)
		adList.lock.RUnlock()
		return
	}
	adList.lock.RUnlock()

	adList.lock.Lock()
	res := mockMySQL(ad)
	adList.cache[ad] = res
	adList.lock.Unlock()

	fmt.Println("success", res)
}

func redisExample() {
	//開啟連線
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer c.Close() //離開時記得關閉

	//清空 reids (測試需求)
	if _, err := c.Do("FLUSHALL"); err != nil {
		fmt.Println(err.Error())
		return
	}

	//嘗試取快取資料
	res, err := redis.String(c.Do("GET", ad))
	if err != nil && !strings.Contains(err.Error(), "nil returned") {
		fmt.Println(err.Error())
		return
	}

	//若無快取更新快取並返回最新值
	if res == "" || err != nil {
		res = mockMySQL(ad)
		if _, err := c.Do("SET", ad, res); err != nil {
			fmt.Println("fail", err.Error())
			return
		}
	}

	fmt.Println("success", res)
}

func mockMySQL(key string) string {
	return `{"ad_name":"九月主打","film_name":"天能","length":168}`
}
