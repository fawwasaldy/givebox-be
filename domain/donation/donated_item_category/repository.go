package donated_item_category

import "context"

type Repository interface {
	Create(ctx context.Context, tx interface{}, donatedItemCategoryEntity DonatedItemCategory) (DonatedItemCategory, error)
}
