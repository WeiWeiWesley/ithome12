package main

import "time"

//FilmModel 系統資料
type FilmModel struct {
	ID          int64      `gorm:"column:id; type: bigint(20); NOT NULL; AUTO_INCREMENT"`
	Name        string     `gorm:"column:name; type: varchar(100); NOT NULL"`
	Category    string     `gorm:"column:category; type: varchar(50); NOT NULL"`
	Length      int64      `gorm:"column:length; type: int(20); NOT NULL"`
	CreatedTime *time.Time `gorm:"column:created_time; type: timestamp; NOT NULL; default:CURRENT_TIMESTAMP"`
}

//TableName users
func (FilmModel) TableName() string {
	return "film"
}
