package models

import (
	"fmt"
	"strings"
	"task5/pkg/util"
	"time"
)

type Mlout struct {
	Id              int64 `gorm:"primary key;AUTO_INCREMENT"`
	TimePoint       time.Time
	StationName     string `gorm:"type:varchar(50)"`
	Substance       string `gorm:"type:varchar(50)"`
	PredictionValue float32
}

type MloutRocket struct {
	TimePoint       string `json:"time_point"`
	PredictionValue string `json:"prediction_value"`
}

//用于发送到rocket mq数据格式定义
type MloutType struct {
	StationName      string        `json:"station_name"`
	Substance        string        `json:"substance"`
	St               int           `json:"st"`                //O3, pm2.5 : 22   湖带 ：32
	PredictionValues []MloutRocket `json:"prediction_values"` // 保存切片
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

func ListPostions() []string {
	util.DaoLog.Info("进行地址查询")
	var mlouts []Mlout
	db.Select("distinct(station_name)").Find(&mlouts)
	var positions []string
	for _, mlout := range mlouts {
		positions = append(positions, mlout.StationName)
	}
	return positions
}

func ListSubstances() []string {
	util.DaoLog.Info("进行污染物查询")
	var mlouts []Mlout
	db.Select("distinct(substance)").Find(&mlouts)
	var substances []string
	for _, mlout := range mlouts {
		substances = append(substances, mlout.Substance)
	}
	return substances
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

func ListMloutTypesBetweenInterval(interval string) []MloutType {
	//返回所有的地点, 污染物以及预测的值
	positions := ListPostions()
	substances := ListSubstances()

	mlouts := ListMloutsBetweenInterval(interval)

	mloutTypes := make([]MloutType, len(positions)*len(substances))
	strs := make([]string, len(positions)*len(substances)) // map的key
	posSubMap := make(map[string]int)                      //映射map

	index := 0
	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(substances); j++ {
			strs[index] = positions[i] + substances[j]
			index++
		}
	}

	//将station_name 和 substance连接作为key， index作为value来做映射map，例如 无锡O3: 0, 无锡PM25: 1...
	for i := 0; i < len(strs); i++ {
		posSubMap[strs[i]] = i
	}

	var timeLayoutStr = "2006-01-02 15:04:05"

	for i := 0; i < len(mlouts); i++ {

		mapKey := mlouts[i].StationName + mlouts[i].Substance // 构造key

		mloutTypes[posSubMap[mapKey]].StationName = mlouts[i].StationName
		mloutTypes[posSubMap[mapKey]].Substance = mlouts[i].Substance
		mloutTypes[posSubMap[mapKey]].St = 22
		// 列表赋值
		mloutTypes[posSubMap[mapKey]].PredictionValues = append(mloutTypes[posSubMap[mapKey]].PredictionValues,
			MloutRocket{TimePoint: mlouts[i].TimePoint.Format(timeLayoutStr),
				PredictionValue: fmt.Sprintf("%.2f", mlouts[i].PredictionValue)})
	}
	return mloutTypes
}
