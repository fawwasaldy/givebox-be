package migration

import (
	"givebox/domain/chat/conversation"
	"givebox/domain/chat/message"
	"givebox/domain/donation/category"
	"givebox/domain/donation/donated_item"
	"givebox/domain/donation/donated_item_category"
	"givebox/domain/donation/image"
	"givebox/domain/profile/profile_review"
	"givebox/infrastructure/database/profile/user"
	"givebox/infrastructure/database/refresh_token"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&user.User{},
		&refresh_token.RefreshToken{},

		&donated_item.DonatedItem{},
		&category.Category{},
		&donated_item_category.DonatedItemCategory{},
		&image.Image{},
		&profile_review.ProfileReview{},
		&conversation.Conversation{},
		&message.Message{},
	); err != nil {
		return err
	}

	return nil
}
