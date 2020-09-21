package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

//ConfigSet 連線設定
type ConfigSet struct {
	Username        string
	Password        string
	Host            string
	DBname          string
	ConnMaxIdel     int
	ConnMaxOpen     int
	ConnMaxLifeTime int64
}

//OpnePool 啟用連線池
func (c *ConfigSet) OpnePool() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", c.Username, c.Password, c.Host, c.DBname))
	if err != nil {
		log.Println(err.Error())
	}

	db.DB().SetMaxIdleConns(c.ConnMaxIdel)
	db.DB().SetMaxOpenConns(c.ConnMaxOpen)
	db.DB().SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second)
	return
}
