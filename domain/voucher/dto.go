package voucher

import "time"

type AvailableVouchersDTO struct {
	Quantity              int
	NextRecoveryDate      time.Time
	AvailableRecoveryInfo []AvailableRecoveryInfoDTO
	AvailableVouchers     []Voucher
}

type AvailableRecoveryInfoDTO struct {
	ReferenceDate  time.Time
	ExpirationDate time.Time
}
