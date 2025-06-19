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
		GetAllDonatedItemsByCityWithPagination(
			ctx context.Context,
			tx interface{},
			city string,
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetDonatedItemByID(ctx context.Context, tx interface{}, id string) (DonatedItem, error)
		CountDonatedItemsByCategoryID(ctx context.Context, tx interface{}, categoryID string) (int64, error)
		Create(ctx context.Context, tx interface{}, donatedItemEntity DonatedItem) (DonatedItem, error)
		Update(ctx context.Context, tx interface{}, donatedItemEntity DonatedItem) (DonatedItem, error)
	}
)
