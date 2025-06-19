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

func (r repository) GetAllProfileReviewsByReceiverIDWithPagination(ctx context.Context, tx interface{}, receiverID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var profileReviewSchemas []Review
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Review{}).
		Where("receiver_id = ?", receiverID)
	if req.Search != "" {
		query = query.Where("message ILIKE ?", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&profileReviewSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	profileReviewEntities := make([]any, len(profileReviewSchemas))
	for i, profileReviewSchema := range profileReviewSchemas {
		profileReviewEntities[i] = SchemaToEntity(profileReviewSchema)
	}
	return pagination.ResponseWithData{
		Data: profileReviewEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) Create(ctx context.Context, tx interface{}, profileReviewEntity review.Review) (review.Review, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return review.Review{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	profileReviewSchema := EntityToSchema(profileReviewEntity)
	if err = db.WithContext(ctx).Create(&profileReviewSchema).Error; err != nil {
		return review.Review{}, err
	}

	profileReviewEntity = SchemaToEntity(profileReviewSchema)
	return profileReviewEntity, nil
}
