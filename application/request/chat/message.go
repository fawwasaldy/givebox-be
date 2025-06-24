package chat

type MessageSend struct {
	ConversationID string `json:"conversation_id" form:"conversation_id" binding:"required"`
	Content        string `json:"content" form:"content" binding:"required"`
}
