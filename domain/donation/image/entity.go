package image

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type Image struct {
	ID            identity.ID
	DonatedItemID identity.ID
	ImageURL      shared.URL
	shared.Timestamp
}
