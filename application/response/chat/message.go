package chat

type Message struct {
	ID               string `json:"id"`
	MessageContent   string `json:"message_content"`
	MessageCreatedAt string `json:"message_created_at"`
	IsMine           bool   `json:"is_mine"`
}

type MessageSend struct {
	ID               string `json:"id"`
	MessageContent   string `json:"message_content"`
	MessageCreatedAt string `json:"message_created_at"`
}
