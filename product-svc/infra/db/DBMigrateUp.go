package db

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"gorm.io/gorm"
)

func PostDBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Category{},
		&domain.Product{},
	)
}
