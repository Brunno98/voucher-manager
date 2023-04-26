package code

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	date1          = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	date2          = date1.AddDate(0, 0, 30)
	subscriptionId = "SOME ID"
)

func TestGetEarliestRecoveryDateNotUsed(t *testing.T) {
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
	availableDates := []time.Time{date1}
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
	tests := map[string]struct {
		AvailableDates []time.Time
		Recovers       []Recovered
		Want           []time.Time
	}{
		"has 1 recover": {
			AvailableDates: []time.Time{date1, date2},
			Recovers:       []Recovered{{ID: 1, CodeId: "SOME CODE", RecoveryDate: date1, ReferenceDate: date1, SubscriptionId: subscriptionId}},
			Want:           []time.Time{date2},
		},
		"has all recover": {
			AvailableDates: []time.Time{date1, date2},
			Recovers: []Recovered{
				{ID: 1, CodeId: "SOME CODE", RecoveryDate: date1, ReferenceDate: date1, SubscriptionId: subscriptionId},
				{ID: 2, CodeId: "SOME OTHER CODE", RecoveryDate: date2, ReferenceDate: date2, SubscriptionId: subscriptionId},
			},
			Want: []time.Time{},
		},
		"has not recover": {
			AvailableDates: []time.Time{date1, date2},
			Recovers:       []Recovered{},
			Want:           []time.Time{date1, date2},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			mockRepository := NewMockRepository(gomock.NewController(t))
			service := NewRecoveryService(mockRepository)

			mockRepository.EXPECT().
				GetRecoveredByReferenceDates(subscriptionId, testCase.AvailableDates).
				Return(testCase.Recovers)

			got := service.RemoveDatesAlreadyRecovered(subscriptionId, testCase.AvailableDates)

			if cmp.Diff(got, testCase.Want) != "" {
				t.Fatalf("Expected: %#v but got: %#v", testCase.Want, got)
			}
		})
	}
}
