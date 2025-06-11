package category

import (
	"context"
	"givebox/domain/donation/category"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) category.Repository {
	return &repository{db: db}
}

func (r repository) GetAllCategories(ctx context.Context, tx interface{}) ([]category.Category, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return nil, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var categorySchemas []Category

	query := db.WithContext(ctx).Model(&Category{})

	if err = query.Find(&categorySchemas).Error; err != nil {
		return nil, err
	}

	categoryEntities := make([]category.Category, len(categorySchemas))
	for i, categorySchema := range categorySchemas {
		categoryEntities[i] = SchemaToEntity(categorySchema)
	}
	return categoryEntities, nil
}
