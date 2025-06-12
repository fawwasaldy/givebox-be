package seeds

import (
	"givebox/infrastructure/database/donation/category"
	"givebox/infrastructure/database/migration/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Category(db *gorm.DB) error {
	hasTable := db.Migrator().HasTable(&category.Category{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&category.Category{}); err != nil {
			return err
		}
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "name"},
		},
		UpdateAll: true,
	}).CreateInBatches(data.Categories, 100).Error
}
