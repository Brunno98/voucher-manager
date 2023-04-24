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

func (service *RecoverServiceImpl) GetEarliestRecoveryDateNotUsed(subscriptionId string, availableDates []time.Time) (earliestDate time.Time, found bool) {
	// (TODO) Impletentar...
	return availableDates[0], true
}

func (service *RecoverServiceImpl) FilterRecoveryAlreadyUsed(dates []time.Time) []time.Time {
	// (TODO) Implementar...
	return dates
}
