package message

import (
	"github.com/google/uuid"
	"givebox/domain/chat/message"
	"time"
)

type Message struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null;column:conversation_id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	Content        string    `gorm:"type:text;not null;column:content"`
	CreatedAt      time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Message) TableName() string {
	return "messages"
}

func EntityToSchema(message message.Message) Message {
	return Message{
		ID:             message.ID.ID,
		ConversationID: message.ConversationID.ID,
		UserID:         message.UserID.ID,
		Content:        message.Content,
		CreatedAt:      message.Timestamp.CreatedAt,
		UpdatedAt:      message.Timestamp.UpdatedAt,
	}
}
