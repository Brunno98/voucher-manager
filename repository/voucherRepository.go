package repository

import (
	"time"

	"github.com/brunno98/voucher-manager/domain/code"
	"github.com/brunno98/voucher-manager/domain/voucher"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	db.AutoMigrate(&voucher.Voucher{})
	return &VoucherRepository{db}
}

func (repository *VoucherRepository) FindByKey(key string) (voucher.Voucher, error) {
	voucher := voucher.Voucher{Key: key}

	if result := repository.db.Find(&voucher); result.Error != nil {
		return voucher, result.Error
	}

	return voucher, nil
}

func (repository *VoucherRepository) Recover(subscriptionId string, referenceDate time.Time, voucher *voucher.Voucher) (code.Code, error) {
	c := code.Code{VoucherId: voucher.ID}
	result := repository.db.
		First(&c).
		Create(&code.Recovered{
			CodeId:         c.Code,
			SubscriptionId: subscriptionId,
			RecoveryDate:   time.Now(),
		})

	if result.Error != nil {
		return c, result.Error
	}

	return c, nil
}

func (repository *VoucherRepository) FindAll() []voucher.Voucher {
	var vouchers []voucher.Voucher
	repository.db.Find(&vouchers)
	return vouchers
}

func (repository *VoucherRepository) PopulateVoucher() {
	if err := repository.db.First(&voucher.Voucher{}).Error; err == gorm.ErrRecordNotFound {
		vouchers := []voucher.Voucher{
			{Key: "foo"}, {Key: "Pokemon"}, {Key: "FreeFire"},
		}
		repository.db.CreateInBatches(vouchers, len(vouchers))
	}
}
