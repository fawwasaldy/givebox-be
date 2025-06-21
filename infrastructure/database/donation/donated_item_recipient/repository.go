package donated_item_recipient

import (
	"context"
	"givebox/domain/donation/donated_item_recipient"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) donated_item_recipient.Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, tx interface{}, donatedItemRecipientEntity donated_item_recipient.DonatedItemRecipient) (donated_item_recipient.DonatedItemRecipient, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item_recipient.DonatedItemRecipient{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	donatedItemRecipientSchema := EntityToSchema(donatedItemRecipientEntity)
	if err = db.WithContext(ctx).Create(&donatedItemRecipientSchema).Error; err != nil {
		return donated_item_recipient.DonatedItemRecipient{}, err
	}

	donatedItemRecipientEntity = SchemaToEntity(donatedItemRecipientSchema)
	return donatedItemRecipientEntity, nil
}

func (r repository) Update(ctx context.Context, tx interface{}, donatedItemRecipientEntity donated_item_recipient.DonatedItemRecipient) (donated_item_recipient.DonatedItemRecipient, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item_recipient.DonatedItemRecipient{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	donatedItemRecipientSchema := EntityToSchema(donatedItemRecipientEntity)
	if err = db.WithContext(ctx).
		Where("donated_item_id = ?", donatedItemRecipientSchema.DonatedItemID).
		Where("recipient_id = ?", donatedItemRecipientSchema.RecipientID).
		Updates(&donatedItemRecipientSchema).Error; err != nil {
		return donated_item_recipient.DonatedItemRecipient{}, err
	}

	donatedItemRecipientEntity = SchemaToEntity(donatedItemRecipientSchema)
	return donatedItemRecipientEntity, nil
}
