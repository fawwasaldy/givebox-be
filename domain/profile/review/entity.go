package review

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type Review struct {
	ID            identity.ID
	DonatedItemID identity.ID
	Message       string
	Rating        shared.LikertScale
	shared.Timestamp
}
