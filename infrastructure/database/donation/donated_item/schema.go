package donated_item

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"time"
)

type DonatedItem struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	DonorID             uuid.UUID `gorm:"type:uuid;not null;column:donor_id"`
	CategoryID          uuid.UUID `gorm:"type:uuid;not null;column:category_id"`
	Status              string    `gorm:"type:varchar(20);not null;column:status"`
	Name                string    `gorm:"type:varchar(50);not null;column:name"`
	Description         string    `gorm:"type:text;column:description"`
	Condition           int       `gorm:"type:int;not null;check:condition >= 1 AND condition <= 5;column:condition"`
	QuantityDescription string    `gorm:"type:text;column:quantity_description"`
	PickCity            string    `gorm:"type:varchar(50);not null;column:pick_city"`
	PickAddress         string    `gorm:"type:varchar(255);not null;column:pick_address"`
	PickingStatus       string    `gorm:"type:varchar(20);not null;column:picking_status"`
	DeliveryTime        string    `gorm:"type:varchar(50);column:delivery_time"`
	IsUrgent            bool      `gorm:"type:boolean;not null;default:false;column:is_urgent"`
	AdditionalNote      string    `gorm:"type:text;column:additional_note"`
	CreatedAt           time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt           time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (DonatedItem) TableName() string {
	return "donated_items"
}

func EntityToSchema(entity donated_item.DonatedItem) DonatedItem {
	return DonatedItem{
		ID:                  entity.ID.ID,
		DonorID:             entity.DonorID.ID,
		CategoryID:          entity.CategoryID.ID,
		Status:              entity.Status.Status,
		Name:                entity.Name,
		Description:         entity.Description,
		Condition:           entity.Condition.Value,
		QuantityDescription: entity.QuantityDescription,
		PickCity:            entity.PickCity,
		PickAddress:         entity.PickAddress,
		PickingStatus:       entity.PickingStatus.Status,
		DeliveryTime:        entity.DeliveryTime,
		IsUrgent:            entity.IsUrgent,
		AdditionalNote:      entity.AdditionalNote,
		CreatedAt:           entity.Timestamp.CreatedAt,
		UpdatedAt:           entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema DonatedItem) donated_item.DonatedItem {
	return donated_item.DonatedItem{
		ID:                  identity.NewIDFromSchema(schema.ID),
		DonorID:             identity.NewIDFromSchema(schema.DonorID),
		CategoryID:          identity.NewIDFromSchema(schema.CategoryID),
		Status:              donated_item.NewStatusFromSchema(schema.Status),
		Name:                schema.Name,
		Description:         schema.Description,
		Condition:           shared.NewLikertScaleFromSchema(schema.Condition),
		QuantityDescription: schema.QuantityDescription,
		PickCity:            schema.PickCity,
		PickAddress:         schema.PickAddress,
		PickingStatus:       donated_item.NewPickingStatusFromSchema(schema.PickingStatus),
		DeliveryTime:        schema.DeliveryTime,
		IsUrgent:            schema.IsUrgent,
		AdditionalNote:      schema.AdditionalNote,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
