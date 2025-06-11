package donated_item_category

import (
	"context"
	"givebox/domain/donation/donated_item_category"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) donated_item_category.Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, tx interface{}, donatedItemCategoryEntity donated_item_category.DonatedItemCategory) (donated_item_category.DonatedItemCategory, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item_category.DonatedItemCategory{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	donatedItemCategorySchema := EntityToSchema(donatedItemCategoryEntity)
	if err = db.WithContext(ctx).Create(&donatedItemCategorySchema).Error; err != nil {
		return donated_item_category.DonatedItemCategory{}, err
	}

	donatedItemCategoryEntity = SchemaToEntity(donatedItemCategorySchema)
	return donatedItemCategoryEntity, nil
}
