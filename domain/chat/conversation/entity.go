package conversation

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type Conversation struct {
	ID                     identity.ID
	DonatedItemRecipientID identity.ID
	LatestMessageID        identity.ID
	shared.Timestamp
}
