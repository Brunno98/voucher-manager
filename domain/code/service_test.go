package code

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
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
				Return(testCase.RecoveredDates, nil).AnyTimes()

			got, err := service.GetEarliestRecoveryDateNotUsed(subscriptionId, testCase.AvailableDates)
			if err != nil {
				t.Fatal(err)
			}
			if got != testCase.Want {
				t.Fatalf("Expected date: %s but got: %s", testCase.Want, got)
			}
		})
	}
}

func TestGetEarliestRecoveryDateNotUsedWithErr(t *testing.T) {
	subscriptionId := "SOME ID"
	availableDates := []time.Time{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)}
	expectedErr := errors.New("some error")
	mockRepository := NewMockRepository(gomock.NewController(t))
	service := NewRecoveryService(mockRepository)
	mockRepository.
		EXPECT().
		GetLastRecoveredDates(subscriptionId, len(availableDates)).
		Return([]time.Time{}, expectedErr).AnyTimes()

	got, err := service.GetEarliestRecoveryDateNotUsed(subscriptionId, availableDates)
	if err != expectedErr {
		t.Fatal(err)
	}
	if got != (time.Time{}) {
		t.Fatalf("Expected date: %s but got: %s", time.Time{}, got)
	}
}

func TestRemoveDatesAlreadyRecovered(t *testing.T) {
	subscriptionId := "SOME SUBSCRIPTION"
	date1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)
	dates := []time.Time{date1, date2}
	recovers := []Recovered{
		{ID: 1, CodeId: "SOME CODE", RecoveryDate: date1, ReferenceDate: date1, SubscriptionId: subscriptionId},
	}
	mockRepository := NewMockRepository(gomock.NewController(t))
	service := NewRecoveryService(mockRepository)

	mockRepository.EXPECT().
		GetRecoveredByReferenceDates(subscriptionId, dates).
		Return(recovers)

	got := service.RemoveDatesAlreadyRecovered(subscriptionId, dates)

	expected := []time.Time{date2}
	if cmp.Diff(got, expected) != "" {
		t.Fatalf("Expected: %#v but got: %#v", expected, got)
	}
}
