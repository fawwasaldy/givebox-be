package image

import (
	"context"
	"givebox/domain/donation/image"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) image.Repository {
	return &repository{db: db}
}

func (r repository) GetAllImagesByDonatedItemID(ctx context.Context, tx interface{}, donatedItemID string) ([]image.Image, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return nil, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var imageSchemas []Image

	query := db.WithContext(ctx).Model(&Image{}).
		Where("donated_item_id = ?", donatedItemID)

	if err = query.Find(&imageSchemas).Error; err != nil {
		return nil, err
	}

	imageEntities := make([]image.Image, len(imageSchemas))
	for i, imageSchema := range imageSchemas {
		imageEntities[i] = SchemaToEntity(imageSchema)
	}
	return imageEntities, nil
}

func (r repository) Create(ctx context.Context, tx interface{}, imageEntity image.Image) (image.Image, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return image.Image{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	imageSchema := EntityToSchema(imageEntity)
	if err = db.WithContext(ctx).Create(&imageSchema).Error; err != nil {
		return image.Image{}, err
	}

	imageEntity = SchemaToEntity(imageSchema)
	return imageEntity, nil
}

func (r repository) Delete(ctx context.Context, tx interface{}, id string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&Image{}).Error; err != nil {
		return err
	}

	return nil
}
