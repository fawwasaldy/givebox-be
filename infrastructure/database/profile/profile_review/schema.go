package profile_review

import (
	"github.com/google/uuid"
	"givebox/domain/identity"
	"givebox/domain/profile/profile_review"
	"givebox/domain/shared"
	"time"
)

type ProfileReview struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	SenderID   uuid.UUID `gorm:"type:uuid;not null;column:sender_id"`
	ReceiverID uuid.UUID `gorm:"type:uuid;not null;column:receiver_id"`
	Message    string    `gorm:"type:varchar(255);not null;column:message"`
	Rating     int       `gorm:"type:int;not null;column:rating"`
	CreatedAt  time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (ProfileReview) TableName() string {
	return "profile_reviews"
}

func EntityToSchema(profileReview profile_review.ProfileReview) ProfileReview {
	return ProfileReview{
		ID:         profileReview.ID.ID,
		SenderID:   profileReview.SenderID.ID,
		ReceiverID: profileReview.ReceiverID.ID,
		Message:    profileReview.Message,
		Rating:     profileReview.Rating.Value,
		CreatedAt:  profileReview.Timestamp.CreatedAt,
		UpdatedAt:  profileReview.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema ProfileReview) profile_review.ProfileReview {
	return profile_review.ProfileReview{
		ID:         identity.NewIDFromSchema(schema.ID),
		SenderID:   identity.NewIDFromSchema(schema.SenderID),
		ReceiverID: identity.NewIDFromSchema(schema.ReceiverID),
		Message:    schema.Message,
		Rating:     shared.NewLikertScaleFromSchema(schema.Rating),
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
