package donated_item_recipient

import "errors"

var (
	ErrorGetDonatedItemRecipientById       = errors.New("failed to get donated item recipient by id")
	ErrorCreateDonatedItemRecipient        = errors.New("failed to create donated item recipient")
	ErrorUpdateDonatedItemRecipient        = errors.New("failed to update donated item recipient")
	ErrorDonatedItemRecipientAlreadyExists = errors.New("donated item recipient already exists")
)
