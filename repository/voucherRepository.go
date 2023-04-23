package repository

import (
	"github.com/brunno98/voucher-manager/domain/voucher"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	db.AutoMigrate(&voucher.Voucher{})
	return &VoucherRepository{db}
}
