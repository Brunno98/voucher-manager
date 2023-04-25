package main

import (
	"log"

	"github.com/brunno98/voucher-manager/domain/code"
	"github.com/brunno98/voucher-manager/domain/voucher"
	"github.com/brunno98/voucher-manager/repository"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("voucher-manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	voucherRepository := repository.NewVoucherRepository(db)
	codeRepository := repository.NewCodeRepository(db)
	recoveredRepository := repository.NewRecoveredRepository(db)

	voucherRepository.PopulateVoucher()
	codeRepository.PopulateCodes()

	recoveryService := code.NewRecoveryService(recoveredRepository)
	voucherService := voucher.NewService(voucherRepository, recoveryService)

	r := gin.Default()

	v1 := r.Group("/api")

	voucher.
		NewVoucherRouter(voucherService).
		InitRouters(v1.Group("voucher"))

	r.Run()
}
