package repository

import (
	"errors"
	"log"
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
	c := code.Code{}
	err := repository.db.
		Raw("SELECT * FROM codes c WHERE c.voucher_id = ? AND c.code NOT IN (SELECT code_id FROM recovereds) LIMIT 1;", voucher.ID).
		Scan(&c).Error
	if err != nil {
		log.Println("ERROR!!! ao recuperar voucher")
		log.Println(err)
		return c, err
	}

	if c.Code == "" {
		return c, errors.New("code not recoverd")
	}

	err = repository.db.Create(&code.Recovered{
		CodeId:         c.Code,
		SubscriptionId: subscriptionId,
		RecoveryDate:   time.Now(),
		ReferenceDate:  referenceDate,
	}).Error
	if err != nil {
		return c, err
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
