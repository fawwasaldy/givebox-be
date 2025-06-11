package message

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type Message struct {
	ID             identity.ID
	ConversationID identity.ID
	UserID         identity.ID
	Content        string
	shared.Timestamp
}
