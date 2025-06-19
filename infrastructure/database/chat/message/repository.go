package message

import (
	"context"
	"givebox/domain/chat/message"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) message.Repository {
	return &repository{db: db}
}

func (r repository) GetAllMessagesByConversationIDWithPagination(ctx context.Context, tx interface{}, conversationID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var messageSchemas []Message
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Message{}).
		Where("conversation_id = ?", conversationID)
	if req.Search != "" {
		query = query.Where("content ILIKE ?", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&messageSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	messageEntities := make([]any, len(messageSchemas))
	for i, messageSchema := range messageSchemas {
		messageEntities[i] = SchemaToEntity(messageSchema)
	}
	return pagination.ResponseWithData{
		Data: messageEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) Create(ctx context.Context, tx interface{}, messageEntity message.Message) (message.Message, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return message.Message{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	messageSchema := EntityToSchema(messageEntity)
	if err = db.WithContext(ctx).Create(&messageSchema).Error; err != nil {
		return message.Message{}, err
	}

	messageEntity = SchemaToEntity(messageSchema)
	return messageEntity, nil
}
