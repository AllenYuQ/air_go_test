package models

import (
	"strings"
	"time"
)

type Mlout struct {
	Id              int64 `gorm:"primary key;AUTO_INCREMENT"`
	TimePoint       time.Time
	StationName     string `gorm:"type:varchar(50)"`
	Substance       string `gorm:"type:varchar(50)"`
	PredictionValue float32
}

func ListMlouts(beginTimePoint string, hour string, substance string) []Mlout {

	const base_format = "2006-01-02 15:04:05"
	beginTime, _ := time.Parse(base_format, beginTimePoint)
	interval := hour + "h"
	duration, _ := time.ParseDuration(interval)
	endTime := beginTime.Add(duration)
	endTimePoint := endTime.Format(base_format)
	var mlouts []Mlout

	db.Where("time_point >= ? AND time_point <= ? AND substance = ?", beginTimePoint, endTimePoint, substance).Find(&mlouts)
	return mlouts
}

func ListMloutsBetweenInterval(interval string) []Mlout {

	nt := time.Now()
	const base_format = "2006-01-02 15:04:05"
	str_time := nt.Format(base_format)
	str_time = strings.Split(str_time, ":")[0] + ":00:00"
	begin_time, _ := time.Parse(base_format, str_time)
	interval = interval + "h"
	duration, _ := time.ParseDuration(interval)
	end_time := begin_time.Add(duration)
	endTimePoint := end_time.Format(base_format)

	var mlouts []Mlout
	db.Where("time_point >= ? AND time_point <= ?", str_time, endTimePoint).Find(&mlouts)
	return mlouts
}
