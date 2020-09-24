package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gorm.io/gorm"
)

//const const
const (
	LOCAL = "local"
)

var (
	closeService = make(chan int, 1)
	closeWait    = 3
	sqlMaster    *gorm.DB
	sqlSlave     *gorm.DB
)

func init() {
	if wait, err := strconv.Atoi(os.Getenv("WAIT")); err == nil {
		closeWait = wait
	}

	//mysql connection pool
	{
		master := &ConfigSet{
			Username:        "root",
			Password:        "root",
			Host:            "localhost:3307",
			DBname:          "netfliiix",
			ConnMaxIdel:     5,
			ConnMaxOpen:     10,
			ConnMaxLifeTime: 300,
		}

		m, err := master.OpnePool()
		if err != nil {
			closeService <- 666
			return
		}
		sqlMaster = m

		slave := &ConfigSet{
			Username:        "root",
			Password:        "root",
			Host:            "localhost:3308",
			DBname:          "netfliiix",
			ConnMaxIdel:     5,
			ConnMaxOpen:     10,
			ConnMaxLifeTime: 300,
		}

		s, err := slave.OpnePool()
		if err != nil {
			closeService <- 666
			return
		}
		sqlSlave = s
	}

}

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	servers := Router()
	for i := range servers {
		server := &http.Server{
			Addr:    servers[i].Port,
			Handler: servers[i].Router,
		}
		go func() {
			err := server.ListenAndServe()
			if err != nil {
				log.Fatalln(err.Error())
				closeService <- 666
				return
			}
		}()
	}

	// 阻塞直到有信號傳入
	select {
	case <-osSignal:
		log.Println("Receive OS exit signal", <-osSignal)

		// 避免未執行完的協程，休息一下再退出
		for t := closeWait; t > 0; t-- {
			log.Printf("Logout after %d second", t)
			time.Sleep(time.Second)
		}

		return
	case code := <-closeService:
		log.Printf("System shutdown! code: %d", code)
		return
	}
}
