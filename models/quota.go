package models

import (
	"strconv"
	"task5/pkg/util"
)

type Quota struct {
	Id     int64 `gorm:"primary key;AUTO_INCREMENT"`
	Count  int
	UserId int64 `gorm:"not null"`
}

func Recharge(id, count string) bool {
	util.DaoLog.Info("进行充值操作")
	user_id, _ := strconv.ParseInt(id, 10, 64)

	var quota Quota
	db.Where(&Quota{UserId: user_id}).First(&quota)

	if quota.Id > 0 {
		count, err := strconv.Atoi(count)
		if err == nil {
			quota.Count = count + quota.Count
			db.Save(&quota)
		} else {
			return false
		}
	} else {
		count, err := strconv.Atoi(count)
		if err == nil {
			quota.UserId = user_id
			quota.Count = count
			db.Save(&quota)
		} else {
			return false
		}
	}
	return true
}

func RechargeHistory(id string) bool {
	util.DaoLog.Info("查询对应的用户是否之前有过充值记录")
	var quota Quota
	db.Select("user_id").Where([]string{id}).First(&quota)
	if quota.Id > 0 {
		return true
	}
	return false
}

func ListQuotaInfo() []Quota {
	var quotas []Quota
	db.Find(&quotas)
	return quotas
}
