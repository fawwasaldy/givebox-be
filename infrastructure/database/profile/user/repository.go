package user

import (
	"context"
	user2 "givebox/domain/profile/user"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) user2.Repository {
	return &repository{db: db}
}

func (r *repository) Register(ctx context.Context, tx interface{}, userEntity user2.User) (user2.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user2.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	userSchema := EntityToSchema(userEntity)
	if err = db.WithContext(ctx).Create(&userSchema).Error; err != nil {
		return user2.User{}, err
	}

	userEntity = SchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *repository) GetUserByID(ctx context.Context, tx interface{}, id string) (user2.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user2.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema User
	if err = db.WithContext(ctx).Where("id = ?", id).Take(&userSchema).Error; err != nil {
		return user2.User{}, err
	}

	userEntity := SchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, tx interface{}, username string) (user2.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user2.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema User
	if err = db.WithContext(ctx).Where("username = ?", username).Take(&userSchema).Error; err != nil {
		return user2.User{}, err
	}

	userEntity := SchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *repository) CheckUsername(ctx context.Context, tx interface{}, username string) (user2.User, bool, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user2.User{}, false, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema User
	if err = db.WithContext(ctx).Where("username = ?", username).Take(&userSchema).Error; err != nil {
		return user2.User{}, false, err
	}

	userEntity := SchemaToEntity(userSchema)
	return userEntity, true, nil
}

func (r *repository) Update(ctx context.Context, tx interface{}, userEntity user2.User) (user2.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user2.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	userSchema := EntityToSchema(userEntity)
	if err = db.WithContext(ctx).Updates(&userSchema).Error; err != nil {
		return user2.User{}, err
	}

	userEntity = SchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *repository) Delete(ctx context.Context, tx interface{}, id string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
