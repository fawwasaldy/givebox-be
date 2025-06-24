package message

const (
	// Failed
	FailedChatToDonor                                  = "Failed to chat to donor user"
	FailedSendMessage                                  = "Failed to send message"
	FailedGetAllConversationsByUserIDWithPagination    = "Failed to get all conversations by user id with pagination"
	FailedGetAllMessagesByConversationIDWithPagination = "Failed to get all messages by conversation id with pagination"

	// Success
	SuccessChatToDonor                                  = "Successfully chat to donor user"
	SuccessSendMessage                                  = "Successfully send message"
	SuccessGetAllConversationsByUserIDWithPagination    = "Successfully get all conversations by user id with pagination"
	SuccessGetAllMessagesByConversationIDWithPagination = "Successfully get all messages by conversation id with pagination"
)
