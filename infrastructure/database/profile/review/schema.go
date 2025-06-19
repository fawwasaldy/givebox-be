package review

import (
	"github.com/google/uuid"
	"givebox/domain/identity"
	"givebox/domain/profile/review"
	"givebox/domain/shared"
	"time"
)

type Review struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	DonatedItemID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_id"`
	Message       string    `gorm:"type:varchar(255);not null;column:message"`
	Rating        int       `gorm:"type:int;not null;chek: rating >= 1 AND rating <= 5;column:rating"`
	CreatedAt     time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Review) TableName() string {
	return "reviews"
}

func EntityToSchema(entity review.Review) Review {
	return Review{
		ID:        entity.ID.ID,
		Message:   entity.Message,
		Rating:    entity.Rating.Value,
		CreatedAt: entity.Timestamp.CreatedAt,
		UpdatedAt: entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema Review) review.Review {
	return review.Review{
		ID:            identity.NewIDFromSchema(schema.ID),
		DonatedItemID: identity.NewIDFromSchema(schema.DonatedItemID),
		Message:       schema.Message,
		Rating:        shared.NewLikertScaleFromSchema(schema.Rating),
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
