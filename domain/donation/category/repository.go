package category

import (
	"context"
)

type (
	Repository interface {
		GetAllCategories(
			ctx context.Context,
			tx interface{},
		) ([]Category, error)
		GetSixCategoriesByMostOpenedDonatedItems(
			ctx context.Context,
			tx interface{},
		) ([]Category, error)
		GetCategoryByID(ctx context.Context, tx interface{}, id string) (Category, error)
	}
)
