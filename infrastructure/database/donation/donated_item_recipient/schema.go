package donated_item_recipient

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item_recipient"
	"givebox/domain/identity"
)

type DonatedItemRecipient struct {
	DonatedItemID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_id"`
	RecipientID   uuid.UUID `gorm:"type:uuid;not null;column:recipient_id"`
}

type Tabler interface {
	TableName() string
}

func (DonatedItemRecipient) TableName() string {
	return "donated_item_recipients"
}

func EntityToSchema(entity donated_item_recipient.DonatedItemRecipient) DonatedItemRecipient {
	return DonatedItemRecipient{
		DonatedItemID: entity.DonatedItemID.ID,
		RecipientID:   entity.RecipientID.ID,
	}
}

func SchemaToEntity(schema DonatedItemRecipient) donated_item_recipient.DonatedItemRecipient {
	return donated_item_recipient.DonatedItemRecipient{
		DonatedItemID: identity.NewIDFromSchema(schema.DonatedItemID),
		RecipientID:   identity.NewIDFromSchema(schema.RecipientID),
	}
}
