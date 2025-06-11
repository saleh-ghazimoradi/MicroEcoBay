package db

import "gorm.io/gorm"

func PostDBMigrateDrop(db *gorm.DB) error {
	err := db.Migrator().DropTable()
	if err != nil {
		return err
	}
	return nil
}
