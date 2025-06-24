package review

import (
	"context"
	"givebox/domain/profile/review"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) review.Repository {
	return &repository{db: db}
}

func (r *repository) GetAllReviewsByDonorIDWithPagination(ctx context.Context, tx interface{}, donorID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var reviewSchemas []Review
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Review{}).
		Joins("JOIN donated_items ON donated_items.id = profile_reviews.donated_item_id").
		Where("donated_items.status = ?", "accepted").
		Where("donated_items.donor_id = ?", donorID)
	if req.Search != "" {
		query = query.Where("message ILIKE ?", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&reviewSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	reviewEntities := make([]any, len(reviewSchemas))
	for i, reviewSchema := range reviewSchemas {
		reviewEntities[i] = SchemaToEntity(reviewSchema)
	}
	return pagination.ResponseWithData{
		Data: reviewEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r *repository) GetAllReviewsByRecipientIDWithPagination(ctx context.Context, tx interface{}, recipientID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var reviewSchemas []Review
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Review{}).
		Joins("JOIN donated_items ON donated_items.id = profile_reviews.donated_item_id").
		Where("donated_items.status = ?", "accepted").
		Joins("JOIN donated_item_recipients ON donated_item_recipients.donated_item_id = donated_items.id").
		Where("donated_item_recipients.recipient_id = ?", recipientID)
	if req.Search != "" {
		query = query.Where("message ILIKE ?", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&reviewSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	reviewEntities := make([]any, len(reviewSchemas))
	for i, reviewSchema := range reviewSchemas {
		reviewEntities[i] = SchemaToEntity(reviewSchema)
	}
	return pagination.ResponseWithData{
		Data: reviewEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r *repository) Create(ctx context.Context, tx interface{}, reviewEntity review.Review) (review.Review, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return review.Review{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	reviewSchema := EntityToSchema(reviewEntity)
	if err = db.WithContext(ctx).Create(&reviewSchema).Error; err != nil {
		return review.Review{}, err
	}

	reviewEntity = SchemaToEntity(reviewSchema)
	return reviewEntity, nil
}
