package db

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"gorm.io/gorm"
)

func PostDBMigrateDrop(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&domain.User{},
		&domain.Address{},
	)
	if err != nil {
		return err
	}
	return nil
}
