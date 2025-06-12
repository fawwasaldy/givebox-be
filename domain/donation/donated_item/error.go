package donated_item

import "errors"

var (
	ErrorOpenDonatedItem         = errors.New("failed to open donated item")
	ErrorRequestDonatedItem      = errors.New("failed to request donated item")
	ErrorAcceptDonatedItem       = errors.New("failed to accept donated item")
	ErrorRejectDonatedItem       = errors.New("failed to reject donated item")
	ErrorTakenDonatedItem        = errors.New("failed to mark donated item as taken")
	ErrorGetAllDonatedItems      = errors.New("failed to get all donated items")
	ErrorGetDonatedItemById      = errors.New("failed to get donated item by id")
	ErrorUpdateDonatedItem       = errors.New("failed to update donated item")
	ErrorDeleteDonatedItem       = errors.New("failed to delete donated item")
	ErrorDonatedItemNotFound     = errors.New("donated item not found")
	ErrorInvalidStatus           = errors.New("invalid status for donated item")
	ErrorInvalidStatusTransition = errors.New("invalid status transition for donated item")
	ErrorInvalidDonatedItemType  = errors.New("failed to convert to donated item type")
)
