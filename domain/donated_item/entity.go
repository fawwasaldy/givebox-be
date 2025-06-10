package donated_item

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type DonatedItem struct {
	ID          identity.ID
	DonorID     identity.ID
	RecipientID identity.ID
	status      Status
	Description string
	Condition   shared.LikertScale
	PickCity    string
	PickAddress string
	shared.Timestamp
}
