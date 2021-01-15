package models

type ErrorInfo struct {
	Id   int8   `gorm:"primary key;AUTO_INCREMENT"`
	Info string `gorm:"type:varchar(255)"`
}
