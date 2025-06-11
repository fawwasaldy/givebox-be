package message

import (
	"context"
	"givebox/platform/pagination"
)

type Repository interface {
	GetAllMessagesByConversationIDWithPagination(
		ctx context.Context,
		tx interface{},
		conversationID string,
		req pagination.Request,
	) (pagination.ResponseWithData, error)
	Create(ctx context.Context, tx interface{}, messageEntity Message) (Message, error)
	Update(ctx context.Context, tx interface{}, messageEntity Message) (Message, error)
	Delete(ctx context.Context, tx interface{}, id string) error
}
