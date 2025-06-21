package category

import "errors"

var (
	ErrorGetAllCategories                         = errors.New("failed to get all categories")
	ErrorGetSixCategoriesByMostOpenedDonatedItems = errors.New("failed to get six categories by most opened donated items")
	ErrorGetCategoryById                          = errors.New("failed to get category by id")
)
