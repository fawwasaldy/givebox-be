package chat

type (
	Conversation struct {
		ID                     string `json:"id"`
		MessageReceiverName    string `json:"message_receiver_name"`
		LatestMessageContent   string `json:"latest_message_content,omitempty"`
		LatestMessageCreatedAt string `json:"latest_message_created_at,omitempty"`
		DonatedItemName        string `json:"donated_item_name"`
	}

	FirstConversation struct {
		DonatedItemID string `json:"donated_item_id"`
		DonorName     string `json:"donor_name"`
		RecipientName string `json:"recipient_name"`
	}
)
