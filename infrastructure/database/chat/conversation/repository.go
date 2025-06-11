package conversation

import (
	"context"
	"givebox/domain/chat/conversation"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) conversation.Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAllConversationsByDonorIDWithPagination(ctx context.Context, tx interface{}, donorID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var conversationSchemas []Conversation
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Conversation{}).
		Joins("JOIN users ON users.id = conversations.recipient_id AND conversations.donor_id = ?", donorID).
		Joins("JOIN messages ON messages.id = conversations.last_message_id")
	if req.Search != "" {
		query = query.Where("users.full_name ILIKE ?", "%"+req.Search+"%")
	}
	query = query.Order("messages.created_at DESC")

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&conversationSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	conversationEntities := make([]any, len(conversationSchemas))
	for i, conversationSchema := range conversationSchemas {
		conversationEntities[i] = SchemaToEntity(conversationSchema)
	}
	return pagination.ResponseWithData{
		Data: conversationEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) GetAllConversationsByRecipientIDWithPagination(ctx context.Context, tx interface{}, recipientID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var conversationSchemas []Conversation
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&Conversation{}).
		Joins("JOIN users ON users.id = conversations.donor_id AND conversations.recipient_id = ?", recipientID).
		Joins("JOIN messages ON messages.id = conversations.last_message_id")
	if req.Search != "" {
		query = query.Where("users.full_name ILIKE ?", "%"+req.Search+"%")
	}
	query = query.Order("messages.created_at DESC")

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&conversationSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	conversationEntities := make([]any, len(conversationSchemas))
	for i, conversationSchema := range conversationSchemas {
		conversationEntities[i] = SchemaToEntity(conversationSchema)
	}
	return pagination.ResponseWithData{
		Data: conversationEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) Create(ctx context.Context, tx interface{}, conversationEntity conversation.Conversation) (conversation.Conversation, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return conversation.Conversation{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	conversationSchema := EntityToSchema(conversationEntity)
	if err = db.WithContext(ctx).Create(&conversationSchema).Error; err != nil {
		return conversation.Conversation{}, err
	}

	conversationEntity = SchemaToEntity(conversationSchema)
	return conversationEntity, nil
}
