package code

import (
	"time"
)

type Code struct {
	Code      string `gorm:"primaryKey"`
	VoucherId uint
}

type Recovered struct {
	ID             uint   `gorm:"primaryKey"`
	CodeId         string `gorm:"not null,unique"`
	RecoveryDate   time.Time
	ReferenceDate  time.Time
	SubscriptionId string `gorm:"not null"`
}
