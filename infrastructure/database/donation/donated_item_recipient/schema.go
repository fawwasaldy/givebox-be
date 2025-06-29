package donated_item_recipient

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item_recipient"
	"givebox/domain/identity"
)

type DonatedItemRecipient struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:id"`
	DonatedItemID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_id"`
	RecipientID   uuid.UUID `gorm:"type:uuid;not null;column:recipient_id"`
	IsAccepted    bool      `gorm:"type:boolean;default:false;not null;column:is_accepted"`
}

type Tabler interface {
	TableName() string
}

func (DonatedItemRecipient) TableName() string {
	return "donated_item_recipients"
}

func EntityToSchema(entity donated_item_recipient.DonatedItemRecipient) DonatedItemRecipient {
	return DonatedItemRecipient{
		ID:            entity.ID.ID,
		DonatedItemID: entity.DonatedItemID.ID,
		RecipientID:   entity.RecipientID.ID,
		IsAccepted:    entity.IsAccepted,
	}
}

func SchemaToEntity(schema DonatedItemRecipient) donated_item_recipient.DonatedItemRecipient {
	return donated_item_recipient.DonatedItemRecipient{
		ID:            identity.NewIDFromSchema(schema.ID),
		DonatedItemID: identity.NewIDFromSchema(schema.DonatedItemID),
		RecipientID:   identity.NewIDFromSchema(schema.RecipientID),
		IsAccepted:    schema.IsAccepted,
	}
}
