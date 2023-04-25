package repository

import (
	"time"

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

func (repository *RecoveredRepository) GetLastRecoveredDates(subscriptionId string, limit int) ([]time.Time, error) {
	// (TODO) implementar...
	return []time.Time{}, nil
}
