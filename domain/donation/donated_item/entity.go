package donated_item

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type DonatedItem struct {
	ID                  identity.ID
	DonorID             identity.ID
	CategoryID          identity.ID
	Status              Status
	Name                string
	Description         string
	Condition           shared.LikertScale
	QuantityDescription string
	PickCity            string
	PickAddress         string
	PickingStatus       PickingStatus
	DeliveryTime        string
	IsUrgent            bool
	AdditionalNote      string
	shared.Timestamp
}
