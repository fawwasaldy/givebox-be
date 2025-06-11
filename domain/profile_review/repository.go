package profile_review

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllProfileReviewsByReceiverIDWithPagination(
		ctx context.Context,
		tx interface{},
		receiverID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	Create(ctx context.Context, tx interface{}, profileReviewEntity ProfileReview) (ProfileReview, error)
}
