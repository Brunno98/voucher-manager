package main

import (
	"log"

	"github.com/brunno98/voucher-manager/repository"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("voucher-manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repository.NewVoucherRepository(db)
	repository.NewCodeRepository(db)
	repository.NewRecoveredRepository(db)

}
