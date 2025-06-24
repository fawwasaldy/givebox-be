package seeds

import (
	"givebox/infrastructure/database/donation/image"
	"givebox/infrastructure/database/migration/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Image seeder
func Image(db *gorm.DB) error {
	images := data.GetImages(db)
	if images == nil {
		return nil
	}

	hasTable := db.Migrator().HasTable(&image.Image{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&image.Image{}); err != nil {
			return err
		}
	}

	return db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(images, 100).Error
}
