package seeds

import (
	"givebox/infrastructure/database/migration/data"
	"givebox/infrastructure/database/profile/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func User(db *gorm.DB) error {
	hasTable := db.Migrator().HasTable(&user.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&user.User{}); err != nil {
			return err
		}
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "email"},
		},
		UpdateAll: true,
	}).CreateInBatches(data.Users, 100).Error
}
