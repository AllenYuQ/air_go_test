package models

import "time"

type QuotaDetail struct {
	Id           int64 `gorm:"primary key;AUTO_INCREMENT"`
	ConsumeTime  time.Time
	ConsumeCount int
	QuotaId      int64 `gorm:"not null"`
}
