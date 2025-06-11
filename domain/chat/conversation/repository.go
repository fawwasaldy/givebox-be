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
	Create(ctx context.Context, tx interface{}, conversationEntity Conversation) (Conversation, error)
}
