package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", c.Username, c.Password, c.Host, c.DBname)
	db, err = gorm.Open(mysql.Open(connInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("OpneMySQLPool：", connInfo, err.Error())
		return nil, err
	}

	return
}
