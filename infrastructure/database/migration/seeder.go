package migration

import (
	"givebox/infrastructure/database/migration/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.Category(db); err != nil {
		return err
	}

	return nil
}
