package donated_item_recipient

import "context"

type Repository interface {
	Create(ctx context.Context, tx interface{}, donatedItemRecipientEntity DonatedItemRecipient) (DonatedItemRecipient, error)
}
