package models

import (
	"task5/pkg/util"
	"time"
)

type Mlmet struct {
	Id            int64 `gorm:"primary key;AUTO_INCREMENT"`
	TimePoint     time.Time
	PositionName  string `gorm:"type:varchar(50)"`
	Precipitation float32
	Temperature   float32
	Ws            float32
	Wd            float32
	Humidity      float32
	Cloudrate     float32
	Skycon        string `gorm:"type:varchar(50)"`
	Pressure      float32
	Visibility    float32
	Dswrf         float32
	Aqi           float32
	Pm25          float32
}

func ListMlmets() []Mlmet {
	util.DaoLog.Info("进行查询")
	var mlmets []Mlmet
	//db.Where("top 10").Find(&mlmets)
	db.Order("time_point desc").Limit(2).Find(&mlmets)
	return mlmets
}
