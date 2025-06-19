package conversation

import (
	"github.com/google/uuid"
	"givebox/domain/chat/conversation"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"time"
)

type Conversation struct {
	ID                     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	DonatedItemRecipientID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_recipient_id"`
	LatestMessageID        uuid.UUID `gorm:"type:uuid;not null;column:latest_message_id"`
	CreatedAt              time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt              time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Conversation) TableName() string {
	return "conversations"
}

func EntityToSchema(entity conversation.Conversation) Conversation {
	return Conversation{
		ID:                     entity.ID.ID,
		DonatedItemRecipientID: entity.DonatedItemRecipientID.ID,
		LatestMessageID:        entity.LatestMessageID.ID,
		CreatedAt:              entity.Timestamp.CreatedAt,
		UpdatedAt:              entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema Conversation) conversation.Conversation {
	return conversation.Conversation{
		ID:                     identity.NewIDFromSchema(schema.ID),
		DonatedItemRecipientID: identity.NewIDFromSchema(schema.DonatedItemRecipientID),
		LatestMessageID:        identity.NewIDFromSchema(schema.LatestMessageID),
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
