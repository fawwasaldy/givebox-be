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

func (r *repository) GetAllConversationsByUserIDWithPagination(ctx context.Context, tx interface{}, userID string, req pagination.Request) (pagination.ResponseWithData, error) {
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
		Joins("JOIN donated_item_recipients ON donated_item_recipients.id = conversations.donated_item_recipient_id").
		Joins("JOIN donated_items ON donated_items.id = donated_item_recipients.donated_item_id").
		Where("donated_item_recipients.recipient_id = ? OR donated_items.donor_id = ?", userID, userID).
		Joins("JOIN messages ON conversations.latest_message_id = messages.id")
	if req.Search != "" {
		query = query.Where("id ILIKE ?", "%"+req.Search+"%")
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

func (r *repository) GetConversationByID(ctx context.Context, tx interface{}, conversationID string) (conversation.Conversation, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return conversation.Conversation{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var conversationSchema Conversation
	if err = db.WithContext(ctx).Where("id = ?", conversationID).Take(&conversationSchema).Error; err != nil {
		return conversation.Conversation{}, err
	}

	conversationEntity := SchemaToEntity(conversationSchema)
	return conversationEntity, nil
}

func (r *repository) Create(ctx context.Context, tx interface{}, conversationEntity conversation.Conversation) (conversation.Conversation, error) {
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

func (r *repository) Update(ctx context.Context, tx interface{}, conversationEntity conversation.Conversation) (conversation.Conversation, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return conversation.Conversation{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	conversationSchema := EntityToSchema(conversationEntity)
	if err = db.WithContext(ctx).Updates(&conversationSchema).Error; err != nil {
		return conversation.Conversation{}, err
	}

	conversationEntity = SchemaToEntity(conversationSchema)
	return conversationEntity, nil
}
