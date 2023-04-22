package code

import (
	"github.com/brunno98/voucher-manager/domain/voucher"
	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	VoucherId voucher.Voucher
}
