package donated_item

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"time"
)

type DonatedItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	DonorID     uuid.UUID `gorm:"type:uuid;not null;column:donor_id"`
	RecipientID uuid.UUID `gorm:"type:uuid;column:recipient_id"`
	Status      string    `gorm:"type:varchar(20);not null;column:status"`
	Name        string    `gorm:"type:varchar(50);not null;column:name"`
	Description string    `gorm:"type:text;column:description"`
	Condition   int       `gorm:"type:int;not null;check:condition >= 0 AND condition <= 5;column:condition"`
	PickCity    string    `gorm:"type:varchar(50);not null;column:pick_city"`
	PickAddress string    `gorm:"type:varchar(255);not null;column:pick_address"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (DonatedItem) TableName() string {
	return "donated_items"
}

func EntityToSchema(entity donated_item.DonatedItem) DonatedItem {
	return DonatedItem{
		ID:          entity.ID.ID,
		DonorID:     entity.DonorID.ID,
		RecipientID: entity.RecipientID.ID,
		Status:      entity.Status.Status,
		Name:        entity.Name,
		Description: entity.Description,
		Condition:   entity.Condition.Value,
		PickCity:    entity.PickCity,
		PickAddress: entity.PickAddress,
		CreatedAt:   entity.Timestamp.CreatedAt,
		UpdatedAt:   entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema DonatedItem) donated_item.DonatedItem {
	return donated_item.DonatedItem{
		ID:          identity.NewIDFromSchema(schema.ID),
		DonorID:     identity.NewIDFromSchema(schema.DonorID),
		RecipientID: identity.NewIDFromSchema(schema.RecipientID),
		Status:      donated_item.NewStatusFromSchema(schema.Status),
		Name:        schema.Name,
		Description: schema.Description,
		Condition:   shared.NewLikertScaleFromSchema(schema.Condition),
		PickCity:    schema.PickCity,
		PickAddress: schema.PickAddress,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
