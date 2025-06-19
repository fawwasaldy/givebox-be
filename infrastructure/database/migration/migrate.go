package migration

import (
	"givebox/infrastructure/database/chat/conversation"
	"givebox/infrastructure/database/chat/message"
	"givebox/infrastructure/database/donation/category"
	"givebox/infrastructure/database/donation/donated_item"
	"givebox/infrastructure/database/donation/donated_item_recipient"
	"givebox/infrastructure/database/donation/image"
	"givebox/infrastructure/database/profile/review"
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
		&donated_item_recipient.DonatedItemRecipient{},
		&image.Image{},
		&review.Review{},
		&conversation.Conversation{},
		&message.Message{},
	); err != nil {
		return err
	}

	return nil
}
