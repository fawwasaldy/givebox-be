package category

import (
	"github.com/google/uuid"
	"givebox/domain/donation/category"
	"givebox/domain/identity"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	Name string    `gorm:"type:varchar(20);not null;uniqueIndex;column:name"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "categories"
}

func EntityToSchema(entity category.Category) Category {
	return Category{
		ID:   entity.ID.ID,
		Name: entity.Name,
	}
}

func SchemaToEntity(schema Category) category.Category {
	return category.Category{
		ID:   identity.NewIDFromSchema(schema.ID),
		Name: schema.Name,
	}
}
