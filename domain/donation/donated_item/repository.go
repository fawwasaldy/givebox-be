package donated_item

import (
	"context"
	"givebox/platform/pagination"
)

type (
	Repository interface {
		GetAllDonatedItemsWithPagination(
			ctx context.Context,
			tx interface{},
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByCategoryIDWithPagination(
			ctx context.Context,
			tx interface{},
			categoryID string,
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByConditionWithPagination(
			ctx context.Context,
			tx interface{},
			conditionScale int,
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByStatusWithPagination(
			ctx context.Context,
			tx interface{},
			status string,
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetAllDonatedItemsBeforeDateWithPagination(
			ctx context.Context,
			tx interface{},
			date string,
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetDonatedItemByID(ctx context.Context, tx interface{}, id string) (DonatedItem, error)
		Create(ctx context.Context, tx interface{}, donatedItemEntity DonatedItem) (DonatedItem, error)
		Update(ctx context.Context, tx interface{}, donatedItemEntity DonatedItem) (DonatedItem, error)
		Delete(ctx context.Context, tx interface{}, id string) error
	}
)
