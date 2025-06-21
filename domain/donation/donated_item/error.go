package donated_item

import "errors"

var (
	ErrorOpenDonatedItem         = errors.New("failed to open donated item")
	ErrorAcceptDonatedItem       = errors.New("failed to accept donated item")
	ErrorGetAllDonatedItems      = errors.New("failed to get all donated items")
	ErrorGetDonatedItemById      = errors.New("failed to get donated item by id")
	ErrorDonatedItemNotFound     = errors.New("donated item not found")
	ErrorInvalidStatus           = errors.New("invalid status for donated item")
	ErrorInvalidPickingStatus    = errors.New("invalid picking status for donated item")
	ErrorInvalidStatusTransition = errors.New("invalid status transition for donated item")
	ErrorInvalidDonatedItemType  = errors.New("failed to convert to donated item type")
)
