package models

import "time"

type Mlpollute struct {
	Id           int64 `gorm:"primary key;AUTO_INCREMENT"`
	TimePoint    time.Time
	Area         string `gorm:"type:varchar(50)"`
	PositionName string `gorm:"type:varchar(50)"`
	Co           string `gorm:"type:varchar(50)"`
	CO_24h       string `gorm:"type:varchar(50);column:co_24h"`
	No2          string `gorm:"type:varchar(50)"`
	No2_24h      string `gorm:"type:varchar(50);column:no2_24h"`
	So2          string `gorm:"type:varchar(50)"`
	So2_24h      string `gorm:"type:varchar(50);column:so2_24"`
	O3           string `gorm:"type:varchar(50)"`
	O3_24h       string `gorm:"type:varchar(50);column:o3_24h"`
	O3_8h_24h    string `gorm:"type:varchar(50);column:o3_8h_24h"`
	Pm10         string `gorm:"type:varchar(50)"`
	Pm10_24h     string `gorm:"type:varchar(50);column:pm10_24h"`
	Pm2_5        string `gorm:"type:varchar(50);column:pm2_5"`
	Pm2_5_24h    string `gorm:"type:varchar(50);column:pm2_5_24h"`
	Quality      string `gorm:"type:varchar(50)"`
}
