package repository

import (
	"github.com/brunno98/voucher-manager/domain/code"
	"gorm.io/gorm"
)

type CodeRepository struct {
	db *gorm.DB
}

func NewCodeRepository(db *gorm.DB) *CodeRepository {
	db.AutoMigrate(&code.Code{})
	return &CodeRepository{db}
}

func (repository *CodeRepository) PopulateCodes() {
	if err := repository.db.First(&code.Code{}).Error; err == gorm.ErrRecordNotFound {
		codes := []code.Code{
			{VoucherId: 1, Code: "b5fdc03e-c2c8-42f8-9415-4f1eb46c6d23"},
			{VoucherId: 1, Code: "c917209c-dd42-43d5-9413-4a74403074a1"},
			{VoucherId: 1, Code: "125d0df8-ce9e-45c4-bce2-727e7d4642e1"},
			{VoucherId: 2, Code: "89aea2ea-d3c0-4517-b536-e52d7e069890"},
			{VoucherId: 2, Code: "ca4b52cb-5cc9-4bc1-8618-98d30b424cfe"},
			{VoucherId: 2, Code: "08fde492-2d22-4087-9e10-a35f4a5583bd"},
			{VoucherId: 3, Code: "b3b86f41-ba74-4a24-a41f-6fcc54b745d1"},
			{VoucherId: 3, Code: "d080c18a-eecc-47cf-b81c-4e7ae207ef37"},
			{VoucherId: 3, Code: "ee3df9cf-b5a5-4c4c-945c-3057387f7863"},
		}
		repository.db.CreateInBatches(&codes, len(codes))
	}
}
