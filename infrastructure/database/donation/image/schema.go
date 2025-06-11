package image

import (
	"github.com/google/uuid"
	"givebox/domain/donation/image"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"time"
)

type Image struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	DonatedItemID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_id"`
	ImageURL      string    `gorm:"type:varchar(255);not null;column:image_url"`
	CreatedAt     time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Image) TableName() string {
	return "images"
}

func EntityToSchema(entity image.Image) Image {
	return Image{
		ID:            entity.ID.ID,
		DonatedItemID: entity.DonatedItemID.ID,
		ImageURL:      entity.ImageURL.Path,
		CreatedAt:     entity.Timestamp.CreatedAt,
		UpdatedAt:     entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema Image) image.Image {
	return image.Image{
		ID:            identity.NewIDFromSchema(schema.ID),
		DonatedItemID: identity.NewIDFromSchema(schema.DonatedItemID),
		ImageURL:      shared.NewURLFromSchema(schema.ImageURL),
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
