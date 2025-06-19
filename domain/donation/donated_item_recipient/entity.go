package donated_item_recipient

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type DonatedItemRecipient struct {
	ID            identity.ID
	DonatedItemID identity.ID
	RecipientID   identity.ID
	shared.Timestamp
}
