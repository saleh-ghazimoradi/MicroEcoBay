package db

import "gorm.io/gorm"

func PostDBMigrator(db *gorm.DB) error {
	return db.AutoMigrate()
}
