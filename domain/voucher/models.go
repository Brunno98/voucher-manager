package voucher

import "github.com/jinzhu/gorm"

type Voucher struct {
	gorm.Model
	Key string
}
