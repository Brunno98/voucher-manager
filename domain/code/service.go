package code

import "time"

type RecoverService interface {
	// Recover(subscriptionId string, referenceDate time.Time, voucherKey string) (Code, error)
	GetEarliestRecoveryDateNotUsed(subscriptionId string, availableDates []time.Time) (earliestDate time.Time, found bool)
	FilterRecoveryAlreadyUsed(dates []time.Time) []time.Time
}

type RecoverServiceImpl struct{}

func NewRecoveryService() *RecoverServiceImpl {
	return &RecoverServiceImpl{}
}
