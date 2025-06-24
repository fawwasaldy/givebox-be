package donated_item_recipient

import "context"

type Repository interface {
	CheckDonatedItemRecipient(ctx context.Context, tx interface{}, donatedItemID, recipientID string) (DonatedItemRecipient, bool, error)
	GetDonatedItemRecipientByID(ctx context.Context, tx interface{}, id string) (DonatedItemRecipient, error)
	Create(ctx context.Context, tx interface{}, donatedItemRecipientEntity DonatedItemRecipient) (DonatedItemRecipient, error)
	Update(ctx context.Context, tx interface{}, donatedItemRecipientEntity DonatedItemRecipient) (DonatedItemRecipient, error)
}
