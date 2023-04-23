package voucher

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetVoucherQuantity(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 60,
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		startDate time.Time
		now       time.Time
		want      int
	}{
		"Same day of start date":      {startDate: startDate, now: startDate, want: 1},
		"same day of new recovery":    {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew), want: 2},
		"one day before new recovery": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew-1), want: 1},
		"one day after new recovery":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew+1), want: 2},
		"one year from startDate":     {startDate: startDate, now: startDate.AddDate(1, 0, 0), want: 13},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetVoucherQuantity(tc.startDate, tc.now)
			if got != tc.want {
				t.Fatalf("Expected: %d but got %d", tc.want, got)
			}
		})
	}
}

func TestGetVoucherQuantityWhenExpiryIsLessthanRenew(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 15,
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		startDate time.Time
		now       time.Time
		want      int
	}{
		"Same day of start date":      {startDate: startDate, now: startDate, want: 1},
		"same day of new recovery":    {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew), want: 2},
		"one day before new recovery": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew-1), want: 1},
		"one day after new recovery":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew+1), want: 2},
		"one year from startDate":     {startDate: startDate, now: startDate.AddDate(1, 0, 0), want: 13},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetVoucherQuantity(tc.startDate, tc.now)
			if got != tc.want {
				t.Fatalf("Expected: %d but got %d", tc.want, got)
			}
		})
	}
}

func TestGetVoucherExpiredQuantity(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 60,
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		startDate time.Time
		now       time.Time
		want      int
	}{
		"Same day of start date":       {startDate: startDate, now: startDate, want: 0},
		"same day of first expiry day": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToExpiry), want: 1},
		"one day after first expiry":   {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToExpiry+1), want: 1},
		"one day before first expiry":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToExpiry-1), want: 0},
		"one year from startDate":      {startDate: startDate, now: startDate.AddDate(1, 0, 0), want: 11},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetVoucherExpiredQuantity(tc.startDate, tc.now)
			if got != tc.want {
				t.Fatalf("Expected: %d but got %d", tc.want, got)
			}
		})
	}
}

func TestGetDateOfAvailableVouchers(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 60,
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		startDate time.Time
		now       time.Time
		want      []time.Time
	}{
		"Same day of start date":         {startDate: startDate, now: startDate, want: []time.Time{startDate}},
		"one day before new recovery":    {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew-1), want: []time.Time{startDate}},
		"same day of new recovery":       {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
		"one day after new recovery":     {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew+1), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
		"one day before second recovery": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2-1), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
		"date of second recovery":        {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2), want: []time.Time{startDate.AddDate(0, 0, r.DaysToRenew), startDate.AddDate(0, 0, r.DaysToRenew*2)}},
		"one day after second recovery":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2+1), want: []time.Time{startDate.AddDate(0, 0, r.DaysToRenew), startDate.AddDate(0, 0, r.DaysToRenew*2)}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetDateOfAvailableVouchers(tc.startDate, tc.now)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestGetNextRecoveryDate(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 60,
	}

	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := map[string]struct {
		startDate time.Time
		now       time.Time
		want      time.Time
	}{
		"Same day of start date":         {startDate: startDate, now: startDate, want: time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)},
		"one day before new recovery":    {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew-1), want: time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)},
		"same day of new recovery":       {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew), want: time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC)},
		"one day after new recovery":     {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew+1), want: time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC)},
		"one day before second recovery": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2-1), want: time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC)},
		"date of second recovery":        {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2), want: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)},
		"one day after second recovery":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2+1), want: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetNextRecoverDate(tc.startDate, tc.now)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}

}

func TestGetExpiryDate(t *testing.T) {
	r := &RecoverParameter{
		DaysToRenew:  30,
		DaysToExpiry: 60,
	}

	tests := map[string]struct {
		startDate time.Time
		want      time.Time
	}{
		"beginning of the month": {startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), want: time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC)},
		"end of the month":       {startDate: time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC), want: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)},
		"middle of the month":    {startDate: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), want: time.Date(2023, 3, 16, 0, 0, 0, 0, time.UTC)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := r.GetExpireDate(tc.startDate)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

// func TestGetDateOfAllPossibleVouchers(t *testing.T) {
// 	r := &RecoverParameter{
// 		DaysToRenew:  30,
// 		DaysToExpiry: 60,
// 	}

// 	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

// 	tests := map[string]struct {
// 		startDate time.Time
// 		now       time.Time
// 		want      []time.Time
// 	}{
// 		"Same day of start date":         {startDate: startDate, now: startDate, want: []time.Time{startDate}},
// 		"one day before new recovery":    {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew-1), want: []time.Time{startDate}},
// 		"same day of new recovery":       {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
// 		"one day after new recovery":     {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew+1), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
// 		"one day before second recovery": {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2-1), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew)}},
// 		"date of second recovery":        {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew), startDate.AddDate(0, 0, r.DaysToRenew*2)}},
// 		"one day after second recovery":  {startDate: startDate, now: startDate.AddDate(0, 0, r.DaysToRenew*2+1), want: []time.Time{startDate, startDate.AddDate(0, 0, r.DaysToRenew), startDate.AddDate(0, 0, r.DaysToRenew*2)}},
// 	}

// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got := r.GetDateOfVouchers(tc.startDate, tc.now)
// 			diff := cmp.Diff(tc.want, got)
// 			if diff != "" {
// 				t.Fatalf(diff)
// 			}
// 		})
// 	}
// }
