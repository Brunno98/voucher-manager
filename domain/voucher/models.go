package voucher

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Voucher struct {
	gorm.Model
	Key string `gorm:"unique"`
}

// Estrutura que represanta a parametrização de resgate.
// Com ela é possível definir o intervalo que um novo resgate é disponibilizado (DaysToRenew) e
// o intervalo que um resgate disponível expira (DaysToExpire).
type RecoverParameter struct {
	DaysToRenew  int
	DaysToExpiry int
}

// Retorna a quantidade de vouchers possíveis entre as datas de inicío e agora,
// desconsiderando expiração.
// Considera que já no dia de ativação é=foi fornecido um voucher.
func (r *RecoverParameter) GetVoucherQuantity(startDate, now time.Time) int {
	deltaDays := now.Sub(startDate).Hours() / 24 // Days
	return int(deltaDays)/r.DaysToRenew + 1
}

// Retorna a quantidade de vouchers expirados entra as datas de inicio e agora.
// Considera que já no dia de ativação foi fornecido um voucher.
func (r *RecoverParameter) GetVoucherExpiredQuantity(startDate, now time.Time) int {
	// O calculo começa a partir da data da 1ª expiração
	dateOfFirstExpiry := startDate.AddDate(0, 0, r.DaysToExpiry)
	if now.Before(dateOfFirstExpiry) {
		// Se a data de agora for anterior a data da primeira expiração
		// então não há vouchers expirados
		return 0
	}
	deltaDays := now.Sub(dateOfFirstExpiry).Hours() / 24 // Days

	// É feita a divisão da diferença de dias pelo dias para um novo resgate
	// para saber quantos resgates foram disponibilizados no intervalo.
	// É somado 1 ao final pois é fornecido um resgate no momento da ativação.
	return (int(deltaDays) / r.DaysToRenew) + 1
}

// Retorna a lista de datas que deram direito à um novo resgate ainda não expirado.
func (r *RecoverParameter) GetDateOfAvailableVouchers(startDate, now time.Time) []time.Time {
	voucherExpired := r.GetVoucherExpiredQuantity(startDate, now)
	quantityVoucher := r.GetVoucherQuantity(startDate, now)

	vouchersDate := []time.Time{}
	for i := voucherExpired; i < quantityVoucher; i++ {
		voucherDate := startDate.AddDate(0, 0, r.DaysToRenew*i)
		vouchersDate = append(vouchersDate, voucherDate)
	}
	return vouchersDate
}

// Retorna a data em que será disponibilizado um novo resgate de voucher
func (r *RecoverParameter) GetNextRecoverDate(startDate, now time.Time) time.Time {
	// Multiplicando a quantidade de voucher disponibilizado em um intervalo pela
	// quantidade de dias para um novo resgate obtem-se a data em que será
	// disponibilizado um novo resgate.
	availableQuantity := r.GetVoucherQuantity(startDate, now)
	return startDate.AddDate(0, 0, r.DaysToRenew*availableQuantity)
}

// Retorna a data em que um resgate expirará
func (r *RecoverParameter) GetExpireDate(availableDate time.Time) time.Time {
	return availableDate.AddDate(0, 0, r.DaysToExpiry)
}
