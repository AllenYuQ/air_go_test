package models

import "time"

type Error struct {
	Id        int64 `gorm:"primary key;AUTO_INCREMENT"`
	Time      time.Time
	ErrorType int8  `gorm:"not null"`
	UserId    int64 `gorm:"not null"`
}
