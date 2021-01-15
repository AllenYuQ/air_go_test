package models

type Quota struct {
	Id     int64 `gorm:"primary key;AUTO_INCREMENT"`
	Count  int
	UserId int64 `gorm:"not null"`
}
