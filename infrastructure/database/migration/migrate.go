package migration

import (
	"givebox/infrastructure/database/refresh_token"
	"givebox/infrastructure/database/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&user.User{},
		&refresh_token.RefreshToken{},
	); err != nil {
		return err
	}

	return nil
}
