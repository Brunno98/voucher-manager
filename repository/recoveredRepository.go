package repository

import (
	"github.com/brunno98/voucher-manager/domain/code"
	"gorm.io/gorm"
)

type RecoveredRepository struct {
	db *gorm.DB
}

func NewRecoveredRepository(db *gorm.DB) *RecoveredRepository {
	db.AutoMigrate(&code.Recovered{})
	return &RecoveredRepository{db}
}
