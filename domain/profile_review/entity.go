package profile_review

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type ProfileReview struct {
	ID         identity.ID
	SenderID   identity.ID
	ReceiverID identity.ID
	Message    string
	Rating     shared.LikertScale
	shared.Timestamp
}
