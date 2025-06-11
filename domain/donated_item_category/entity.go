package donated_item_category

import "givebox/domain/identity"

type DonatedItemCategory struct {
	DonatedItemID identity.ID
	CategoryID    identity.ID
}
