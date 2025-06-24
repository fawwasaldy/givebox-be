package migration

import (
	"givebox/infrastructure/database/migration/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	//if err := seeds.Category(db); err != nil {
	//	return err
	//}

	if err := seeds.User(db); err != nil {
		return err
	}

	if err := seeds.DonatedItem(db); err != nil {
		return err
	}

	if err := seeds.Image(db); err != nil {
		return err
	}

	return nil
}
