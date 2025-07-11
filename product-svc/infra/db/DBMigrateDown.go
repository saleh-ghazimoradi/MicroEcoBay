package db

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"gorm.io/gorm"
)

func PostDBMigrateDrop(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&domain.Category{},
		&domain.Product{},
	)
	if err != nil {
		return err
	}
	return nil
}
