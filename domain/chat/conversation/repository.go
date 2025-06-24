package conversation

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllConversationsByUserIDWithPagination(
		ctx context.Context,
		tx interface{},
		userID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	GetConversationByID(ctx context.Context, tx interface{}, conversationID string) (Conversation, error)
	Create(ctx context.Context, tx interface{}, conversationEntity Conversation) (Conversation, error)
	Update(ctx context.Context, tx interface{}, conversationEntity Conversation) (Conversation, error)
}
