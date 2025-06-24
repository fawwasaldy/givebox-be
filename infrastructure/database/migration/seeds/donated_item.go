package seeds

import (
	"givebox/infrastructure/database/donation/donated_item"
	"givebox/infrastructure/database/migration/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DonatedItem(db *gorm.DB) error {
	donatedItems := data.GetDonatedItems(db)

	hasTable := db.Migrator().HasTable(&donated_item.DonatedItem{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&donated_item.DonatedItem{}); err != nil {
			return err
		}
	}

	return db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(donatedItems, 100).Error
}
