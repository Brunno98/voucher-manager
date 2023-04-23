package voucher

import (
	"time"

	"github.com/brunno98/voucher-manager/domain/code"
)

type Repository interface {
	FindByKey(key string) (Voucher, error)
	Recover(subscriptionId string, referenceDate time.Time, voucher *Voucher) (code.Code, error)
	FindAll() []Voucher
}
