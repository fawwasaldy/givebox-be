package donated_item_category

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item_category"
	"givebox/domain/identity"
)

type DonatedItemCategory struct {
	DonatedItemID uuid.UUID `gorm:"type:uuid;not null;column:donated_item_id"`
	CategoryID    uuid.UUID `gorm:"type:uuid;not null;column:category_id"`
}

type Tabler interface {
	TableName() string
}

func (DonatedItemCategory) TableName() string {
	return "donated_items_categories"
}

func EntityToSchema(donatedItemCategory donated_item_category.DonatedItemCategory) DonatedItemCategory {
	return DonatedItemCategory{
		DonatedItemID: donatedItemCategory.DonatedItemID.ID,
		CategoryID:    donatedItemCategory.CategoryID.ID,
	}
}

func SchemaToEntity(schema DonatedItemCategory) donated_item_category.DonatedItemCategory {
	return donated_item_category.DonatedItemCategory{
		DonatedItemID: identity.NewIDFromSchema(schema.DonatedItemID),
		CategoryID:    identity.NewIDFromSchema(schema.CategoryID),
	}
}
