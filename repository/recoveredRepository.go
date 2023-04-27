package repository

import (
	"log"
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
	recovers := []code.Recovered{}
	err := repository.db.
		Limit(limit).
		Order("reference_date DESC").
		Where(&code.Recovered{SubscriptionId: subscriptionId}).
		Find(&recovers).Error
	if err != nil {
		return []time.Time{}, err
	}

	recoveredDates := []time.Time{}
	for _, r := range recovers {
		recoveredDates = append(recoveredDates, r.ReferenceDate)
	}

	return recoveredDates, nil
}

func (repository *RecoveredRepository) GetRecoveredByReferenceDates(subscriptionId string, dates []time.Time) []code.Recovered {
	recovers := []code.Recovered{}

	err := repository.db.
		Where(map[string]interface{}{
			"reference_date":  dates,
			"subscription_id": subscriptionId,
		}).
		Find(&recovers).Error
	if err != nil {
		log.Fatal(err)
		return []code.Recovered{}
	}

	return recovers
}
