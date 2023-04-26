package code

import (
	"time"
)

type Code struct {
	Code      string `gorm:"primaryKey"`
	VoucherId uint
}

type Recovered struct {
	ID             uint
	CodeId         string
	RecoveryDate   time.Time
	ReferenceDate  time.Time
	SubscriptionId string
}
