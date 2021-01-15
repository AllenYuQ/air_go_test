package models

type Role struct {
	id   int8   `gorm:"primary key;AUTO_INCREMENT"`
	Role string `gorm:"type:varchar(20)"`
}
