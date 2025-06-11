package conversation

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllConversationsByDonorIDWithPagination(
		ctx context.Context,
		tx interface{},
		donorID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	GetAllConversationsByRecipientIDWithPagination(
		ctx context.Context,
		tx interface{},
		recipientID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	Create(ctx context.Context, tx interface{}, conversationEntity Conversation) (Conversation, error)
}
