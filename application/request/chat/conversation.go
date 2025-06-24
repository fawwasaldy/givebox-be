package chat

type FirstConversation struct {
	DonatedItemID string `json:"donated_item_id" form:"donated_item_id" binding:"required"`
}
