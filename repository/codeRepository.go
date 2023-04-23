package repository

import (
	"github.com/brunno98/voucher-manager/domain/code"
	"gorm.io/gorm"
)

type CodeRepository struct {
	db *gorm.DB
}

func NewCodeRepository(db *gorm.DB) *CodeRepository {
	db.AutoMigrate(&code.Code{})
	return &CodeRepository{db}
}
