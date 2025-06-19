package conversation

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllConversationsWithPagination(
		ctx context.Context,
		tx interface{},
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	Create(ctx context.Context, tx interface{}, conversationEntity Conversation) (Conversation, error)
}
