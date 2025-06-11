package image

import "context"

type (
	Repository interface {
		GetAllImagesByDonatedItemID(
			ctx context.Context,
			tx interface{},
			donatedItemID string,
		) ([]Image, error)
		Create(ctx context.Context, tx interface{}, imageEntity Image) (Image, error)
		Delete(ctx context.Context, tx interface{}, id string) error
	}
)
