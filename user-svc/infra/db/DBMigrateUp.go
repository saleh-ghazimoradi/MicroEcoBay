package db

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"gorm.io/gorm"
)

func PostDBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
	)
}
