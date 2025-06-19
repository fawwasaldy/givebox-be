package review

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllReviewsByDonorIDWithPagination(
		ctx context.Context,
		tx interface{},
		donorID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	GetAllReviewsByRecipientIDWithPagination(
		ctx context.Context,
		tx interface{},
		recipientID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	Create(ctx context.Context, tx interface{}, profileReviewEntity Review) (Review, error)
}
