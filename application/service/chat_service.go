package service

import (
	"context"
	"errors"
	"givebox/application"
	request_chat "givebox/application/request/chat"
	response_chat "givebox/application/response/chat"
	"givebox/domain/chat/conversation"
	"givebox/domain/chat/message"
	"givebox/domain/donation/donated_item"
	"givebox/domain/donation/donated_item_recipient"
	"givebox/domain/identity"
	"givebox/domain/profile/user"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
	"gorm.io/gorm"
)

type (
	ChatService interface {
		ChatToDonor(ctx context.Context, recipientID string, req request_chat.FirstConversation) (response_chat.FirstConversation, error)
		SendMessage(ctx context.Context, senderID string, req request_chat.MessageSend) (response_chat.MessageSend, error)
		GetAllConversationsByUserIDWithPagination(ctx context.Context, userID string, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllMessagesByConversationIDWithPagination(ctx context.Context, conversationID string, myUserID string, req pagination.Request) (pagination.ResponseWithData, error)
	}

	chatService struct {
		conversationRepository         conversation.Repository
		messageRepository              message.Repository
		donatedItemRecipientRepository donated_item_recipient.Repository
		donatedItemRepository          donated_item.Repository
		userRepository                 user.Repository
		transaction                    interface{}
	}
)

func NewChatService(
	conversationRepository conversation.Repository,
	messageRepository message.Repository,
	donatedItemRecipientRepository donated_item_recipient.Repository,
	donatedItemRepository donated_item.Repository,
	userRepository user.Repository,
	transaction interface{},
) ChatService {
	return &chatService{
		conversationRepository:         conversationRepository,
		messageRepository:              messageRepository,
		donatedItemRecipientRepository: donatedItemRecipientRepository,
		donatedItemRepository:          donatedItemRepository,
		userRepository:                 userRepository,
		transaction:                    transaction,
	}
}

func (s *chatService) ChatToDonor(ctx context.Context, recipientID string, req request_chat.FirstConversation) (response_chat.FirstConversation, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedDonatedItem, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.DonatedItemID)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}
	retrievedRecipient, err := s.userRepository.GetUserByID(ctx, tx, recipientID)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	_, flag, err := s.donatedItemRecipientRepository.CheckDonatedItemRecipient(ctx, tx, retrievedDonatedItem.ID.String(), retrievedRecipient.ID.String())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response_chat.FirstConversation{}, err
	}

	if flag {
		return response_chat.FirstConversation{}, donated_item_recipient.ErrorDonatedItemRecipientAlreadyExists
	}

	donatedItemRecipientEntity := donated_item_recipient.DonatedItemRecipient{
		DonatedItemID: retrievedDonatedItem.ID,
		RecipientID:   retrievedRecipient.ID,
	}

	createdDonatedItemRecipient, err := s.donatedItemRecipientRepository.Create(ctx, tx, donatedItemRecipientEntity)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	conversationEntity := conversation.Conversation{
		DonatedItemRecipientID: createdDonatedItemRecipient.ID,
	}

	_, err = s.conversationRepository.Create(ctx, tx, conversationEntity)
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	retrievedDonor, err := s.userRepository.GetUserByID(ctx, tx, retrievedDonatedItem.DonorID.String())
	if err != nil {
		return response_chat.FirstConversation{}, err
	}

	return response_chat.FirstConversation{
		DonatedItemID: createdDonatedItemRecipient.DonatedItemID.String(),
		DonorName:     retrievedDonor.Name.FullName(),
		RecipientName: retrievedRecipient.Name.FullName(),
	}, nil
}

func (s *chatService) SendMessage(ctx context.Context, senderID string, req request_chat.MessageSend) (response_chat.MessageSend, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_chat.MessageSend{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_chat.MessageSend{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	conversationID := identity.NewID(req.ConversationID)
	userID := identity.NewID(senderID)

	messageEntity := message.Message{
		ConversationID: conversationID,
		UserID:         userID,
		Content:        req.Content,
	}

	createdMessage, err := s.messageRepository.Create(ctx, tx, messageEntity)
	if err != nil {
		return response_chat.MessageSend{}, err
	}

	conversationEntity := conversation.Conversation{
		ID:              createdMessage.ConversationID,
		LatestMessageID: createdMessage.ID,
	}

	_, err = s.conversationRepository.Update(ctx, tx, conversationEntity)

	return response_chat.MessageSend{
		ID:               createdMessage.ID.String(),
		MessageContent:   createdMessage.Content,
		MessageCreatedAt: createdMessage.CreatedAt.String(),
	}, nil
}

func (s *chatService) GetAllConversationsByUserIDWithPagination(ctx context.Context, userID string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.conversationRepository.GetAllConversationsByUserIDWithPagination(ctx, s.transaction, userID, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedConversations := range retrievedData.Data {
		conversationEntity, ok := retrievedConversations.(conversation.Conversation)
		if !ok {
			return pagination.ResponseWithData{}, conversation.ErrorInvalidConversationType
		}

		var receiverID string

		var retrievedDonatedItemRecipient donated_item_recipient.DonatedItemRecipient
		retrievedDonatedItemRecipient, err = s.donatedItemRecipientRepository.GetDonatedItemRecipientByID(ctx, s.transaction, conversationEntity.DonatedItemRecipientID.String())
		if err != nil {
			return pagination.ResponseWithData{}, donated_item_recipient.ErrorGetDonatedItemRecipientById
		}

		if retrievedDonatedItemRecipient.RecipientID.String() != userID {
			receiverID = retrievedDonatedItemRecipient.RecipientID.String()
		}

		var retrievedDonatedItem donated_item.DonatedItem
		retrievedDonatedItem, err = s.donatedItemRepository.GetDonatedItemByID(ctx, s.transaction, retrievedDonatedItemRecipient.DonatedItemID.String())
		if err != nil {
			return pagination.ResponseWithData{}, donated_item.ErrorGetDonatedItemById
		}

		if retrievedDonatedItem.DonorID.String() != userID {
			receiverID = retrievedDonatedItem.DonorID.String()
		}

		var retrievedReceiver user.User
		retrievedReceiver, err = s.userRepository.GetUserByID(ctx, s.transaction, receiverID)
		if err != nil {
			return pagination.ResponseWithData{}, user.ErrorGetUserById
		}

		var retrievedMessage message.Message
		if conversationEntity.LatestMessageID.ID == identity.NilID {
			retrievedMessage = message.Message{
				ID:             identity.NewID(identity.NilID.String()),
				ConversationID: identity.NewID(identity.NilID.String()),
				UserID:         identity.NewID(identity.NilID.String()),
				Content:        "",
			}
		} else {
			retrievedMessage, err = s.messageRepository.GetMessageByID(ctx, s.transaction, conversationEntity.LatestMessageID.String())
			if err != nil {
				return pagination.ResponseWithData{}, message.ErrorGetMessageById
			}
		}

		data = append(data, response_chat.Conversation{
			ID:                     conversationEntity.ID.String(),
			MessageReceiverName:    retrievedReceiver.Name.FullName(),
			LatestMessageContent:   retrievedMessage.Content,
			LatestMessageCreatedAt: retrievedMessage.CreatedAt.String(),
			DonatedItemName:        retrievedDonatedItem.Name,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *chatService) GetAllMessagesByConversationIDWithPagination(ctx context.Context, conversationID string, myUserID string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.messageRepository.GetAllMessagesByConversationIDWithPagination(ctx, s.transaction, conversationID, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedMessages := range retrievedData.Data {
		messageEntity, ok := retrievedMessages.(message.Message)
		if !ok {
			return pagination.ResponseWithData{}, message.ErrorInvalidMessageType
		}

		var isMine bool
		if messageEntity.UserID.String() == myUserID {
			isMine = true
		} else {
			isMine = false
		}

		data = append(data, response_chat.Message{
			ID:               messageEntity.ID.String(),
			MessageContent:   messageEntity.Content,
			MessageCreatedAt: messageEntity.CreatedAt.String(),
			IsMine:           isMine,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}
