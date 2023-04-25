package code

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestGetEarliestRecoveryDateNotUsed(t *testing.T) {
	date1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := date1.AddDate(0, 0, 30)
	subscriptionId := "SOME ID"

	tests := map[string]struct {
		AvailableDates []time.Time
		RecoveredDates []time.Time
		Want           time.Time
	}{
		"não tem resgate usado":                 {AvailableDates: []time.Time{date1, date2}, RecoveredDates: []time.Time{}, Want: date1},
		"resgate não esta nas data disponiveis": {AvailableDates: []time.Time{date1, date2}, RecoveredDates: []time.Time{date1.AddDate(0, 0, -30)}, Want: date1},
		"1 resgate está nas datas disponiveis":  {AvailableDates: []time.Time{date1, date2}, RecoveredDates: []time.Time{date1}, Want: date2},
		"todas as datas já foram resgatadas":    {AvailableDates: []time.Time{date1, date2}, RecoveredDates: []time.Time{date1, date2}, Want: time.Time{}},
		"não tem resgate disponivel":            {AvailableDates: []time.Time{}, RecoveredDates: []time.Time{}, Want: time.Time{}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			mockRepository := NewMockRepository(gomock.NewController(t))
			service := NewRecoveryService(mockRepository)
			mockRepository.
				EXPECT().
				GetLastRecoveredDates(subscriptionId, len(testCase.AvailableDates)).
				Return(testCase.RecoveredDates).AnyTimes()

			got := service.GetEarliestRecoveryDateNotUsed(subscriptionId, testCase.AvailableDates)
			if got != testCase.Want {
				t.Fatalf("Expected date: %s but got: %s", testCase.Want, got)
			}
		})
	}
}
