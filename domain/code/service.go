package code

import (
	"time"
)

type RecoverService interface {
	// Recover(subscriptionId string, referenceDate time.Time, voucherKey string) (Code, error)
	GetEarliestRecoveryDateNotUsed(subscriptionId string, availableDates []time.Time) (time.Time, error)
	RemoveDatesAlreadyRecovered(subscriptionId string, dates []time.Time) []time.Time
}

type RecoverServiceImpl struct {
	Repository
}

func NewRecoveryService(r Repository) *RecoverServiceImpl {
	return &RecoverServiceImpl{r}
}

// Dada a lista de datas de resgates de um usuário, é rotornado a data não usada mais antiga dessa lista.
// availableDates deve estar ordenado começando da data mais antiga.
func (service *RecoverServiceImpl) GetEarliestRecoveryDateNotUsed(subscriptionId string, availableDates []time.Time) (time.Time, error) {
	if len(availableDates) == 0 {
		return time.Time{}, nil
	}

	recoveredDates, err := service.Repository.GetLastRecoveredDates(subscriptionId, len(availableDates))
	if err != nil {
		return time.Time{}, err
	}
	if len(recoveredDates) == 0 {
		return availableDates[0], nil
	}

	var validDate time.Time
	for _, available := range availableDates {
		isRecovered := false
		for _, recovered := range recoveredDates {
			if available == recovered {
				isRecovered = true
				break
			}
		}
		if isRecovered {
			continue
		}
		validDate = available
		break
	}

	return validDate, nil
}

func (service *RecoverServiceImpl) RemoveDatesAlreadyRecovered(subscriptionId string, dates []time.Time) []time.Time {
	recovers := service.Repository.GetRecoveredByReferenceDates(subscriptionId, dates)

	availables := []time.Time{}
	for _, date := range dates {
		isRecovered := false
		for _, recovered := range recovers {
			if date == recovered.ReferenceDate {
				isRecovered = true
				break
			}
		}
		if !isRecovered {
			availables = append(availables, date)
		}
	}

	return availables
}
