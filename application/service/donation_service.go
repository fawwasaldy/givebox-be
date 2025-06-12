package service

import (
	"context"
	"givebox/application"
	request_donation "givebox/application/request/donation"
	response_donation "givebox/application/response/donation"
	"givebox/domain/donation/donated_item"
	"givebox/domain/donation/image"
	"givebox/domain/identity"
	"givebox/domain/shared"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type (
	DonationService interface {
		GetAllDonatedItemsWithPagination(ctx context.Context, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByCategoryIDWithPagination(ctx context.Context, categoryID string, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByConditionWithPagination(ctx context.Context, condition int, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByStatusWithPagination(ctx context.Context, status string, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsBeforeDateWithPagination(ctx context.Context, date string, req pagination.Request) (pagination.ResponseWithData, error)
		GetDonatedItemByID(ctx context.Context, id string) (response_donation.DonationItem, error)
		OpenDonatedItem(ctx context.Context, donorID string, req request_donation.DonationItemOpen) (response_donation.DonationItemOpen, error)
		RequestDonatedItem(ctx context.Context, recipientID string, req request_donation.DonationItemRequest) (response_donation.DonationItemRequest, error)
		AcceptDonatedItem(ctx context.Context, req request_donation.DonationItemAccept) (response_donation.DonationItemAccept, error)
		RejectDonatedItem(ctx context.Context, req request_donation.DonationItemReject) (response_donation.DonationItemReject, error)
		TakenDonatedItem(ctx context.Context, req request_donation.DonationItemTaken) (response_donation.DonationItemTaken, error)
		UpdateDonatedItem(ctx context.Context, req request_donation.DonationItemUpdate) (response_donation.DonationItemUpdate, error)
		DeleteDonatedItem(ctx context.Context, id string) error
	}

	donationService struct {
		donatedItemRepository donated_item.Repository
		imageRepository       image.Repository
		transaction           interface{}
	}
)

func NewDonationService(
	donatedItemRepository donated_item.Repository,
	imageRepository image.Repository,
	transaction interface{},
) DonationService {
	return &donationService{
		donatedItemRepository: donatedItemRepository,
		imageRepository:       imageRepository,
		transaction:           transaction,
	}
}

func (s *donationService) GetAllDonatedItemsWithPagination(ctx context.Context, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsWithPagination(ctx, nil, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		data = append(data, response_donation.DonationItem{
			ID:          donatedItemEntity.ID.String(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetAllDonatedItemsByCategoryIDWithPagination(ctx context.Context, categoryID string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsByCategoryIDWithPagination(ctx, nil, categoryID, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		data = append(data, response_donation.DonationItem{
			ID:          donatedItemEntity.ID.String(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetAllDonatedItemsByConditionWithPagination(ctx context.Context, condition int, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsByConditionWithPagination(ctx, nil, condition, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		data = append(data, response_donation.DonationItem{
			ID:          donatedItemEntity.ID.String(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetAllDonatedItemsByStatusWithPagination(ctx context.Context, status string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsByStatusWithPagination(ctx, nil, status, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		data = append(data, response_donation.DonationItem{
			ID:          donatedItemEntity.ID.String(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetAllDonatedItemsBeforeDateWithPagination(ctx context.Context, date string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsBeforeDateWithPagination(ctx, nil, date, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		data = append(data, response_donation.DonationItem{
			ID:          donatedItemEntity.ID.String(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetDonatedItemByID(ctx context.Context, id string) (response_donation.DonationItem, error) {
	retrievedDonatedItem, err := s.donatedItemRepository.GetDonatedItemByID(ctx, nil, id)
	if err != nil {
		return response_donation.DonationItem{}, donated_item.ErrorGetDonatedItemById
	}

	return response_donation.DonationItem{
		ID:          retrievedDonatedItem.ID.String(),
		Name:        retrievedDonatedItem.Name,
		Description: retrievedDonatedItem.Description,
		Condition:   retrievedDonatedItem.Condition.Value,
		PickCity:    retrievedDonatedItem.PickCity,
	}, nil
}

func (s *donationService) OpenDonatedItem(ctx context.Context, donorID string, req request_donation.DonationItemOpen) (response_donation.DonationItemOpen, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonationItemOpen{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonationItemOpen{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	condition, err := shared.NewLikertScale(req.Condition)
	if err != nil {
		return response_donation.DonationItemOpen{}, donated_item.ErrorOpenDonatedItem
	}
	status, err := donated_item.NewStatus(donated_item.StatusOpened)
	if err != nil {
		return response_donation.DonationItemOpen{}, donated_item.ErrorInvalidStatus
	}

	donatedItemEntity := donated_item.DonatedItem{
		DonorID:     identity.NewID(donorID),
		Status:      status,
		Name:        req.Name,
		Description: req.Description,
		Condition:   condition,
		PickCity:    req.PickCity,
		PickAddress: req.PickAddress,
	}

	createdDonatedItem, err := s.donatedItemRepository.Create(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemOpen{}, donated_item.ErrorOpenDonatedItem
	}

	for _, reqImage := range req.Images {
		var imageURL shared.URL
		imageURL, err = shared.NewURL(reqImage)
		if err != nil {
			return response_donation.DonationItemOpen{}, donated_item.ErrorOpenDonatedItem
		}
		imageEntity := image.Image{
			DonatedItemID: createdDonatedItem.ID,
			ImageURL:      imageURL,
		}

		if _, err = s.imageRepository.Create(ctx, tx, imageEntity); err != nil {
			return response_donation.DonationItemOpen{}, donated_item.ErrorOpenDonatedItem
		}
	}

	return response_donation.DonationItemOpen{
		ID:          createdDonatedItem.ID.String(),
		DonorID:     createdDonatedItem.DonorID.String(),
		Status:      createdDonatedItem.Status.Status,
		Name:        createdDonatedItem.Name,
		Description: createdDonatedItem.Description,
		Condition:   createdDonatedItem.Condition.Value,
		PickCity:    createdDonatedItem.PickCity,
		PickAddress: createdDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) RequestDonatedItem(ctx context.Context, recipientID string, req request_donation.DonationItemRequest) (response_donation.DonationItemRequest, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonationItemRequest{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonationItemRequest{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	donatedItemEntity, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.ID)
	if err != nil {
		return response_donation.DonationItemRequest{}, donated_item.ErrorDonatedItemNotFound
	}

	if donatedItemEntity.Status.Status != donated_item.StatusOpened {
		return response_donation.DonationItemRequest{}, donated_item.ErrorInvalidStatusTransition
	}

	donatedItemEntity.RecipientID = identity.NewID(recipientID)
	donatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusRequested)
	if err != nil {
		return response_donation.DonationItemRequest{}, donated_item.ErrorInvalidStatus
	}

	updatedDonatedItem, err := s.donatedItemRepository.Update(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemRequest{}, donated_item.ErrorRequestDonatedItem
	}

	return response_donation.DonationItemRequest{
		ID:          updatedDonatedItem.ID.String(),
		DonorID:     updatedDonatedItem.DonorID.String(),
		RecipientID: updatedDonatedItem.RecipientID.String(),
		Status:      updatedDonatedItem.Status.Status,
		Name:        updatedDonatedItem.Name,
		Description: updatedDonatedItem.Description,
		Condition:   updatedDonatedItem.Condition.Value,
		PickCity:    updatedDonatedItem.PickCity,
		PickAddress: updatedDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) AcceptDonatedItem(ctx context.Context, req request_donation.DonationItemAccept) (response_donation.DonationItemAccept, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonationItemAccept{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonationItemAccept{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	donatedItemEntity, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.ID)
	if err != nil {
		return response_donation.DonationItemAccept{}, donated_item.ErrorDonatedItemNotFound
	}

	if donatedItemEntity.Status.Status != donated_item.StatusRequested {
		return response_donation.DonationItemAccept{}, donated_item.ErrorInvalidStatusTransition
	}

	donatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusAccepted)
	if err != nil {
		return response_donation.DonationItemAccept{}, donated_item.ErrorInvalidStatus
	}

	updatedDonatedItem, err := s.donatedItemRepository.Update(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemAccept{}, donated_item.ErrorAcceptDonatedItem
	}

	return response_donation.DonationItemAccept{
		ID:          updatedDonatedItem.ID.String(),
		DonorID:     updatedDonatedItem.DonorID.String(),
		RecipientID: updatedDonatedItem.RecipientID.String(),
		Status:      updatedDonatedItem.Status.Status,
		Name:        updatedDonatedItem.Name,
		Description: updatedDonatedItem.Description,
		Condition:   updatedDonatedItem.Condition.Value,
		PickCity:    updatedDonatedItem.PickCity,
		PickAddress: updatedDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) RejectDonatedItem(ctx context.Context, req request_donation.DonationItemReject) (response_donation.DonationItemReject, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonationItemReject{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonationItemReject{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	donatedItemEntity, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.ID)
	if err != nil {
		return response_donation.DonationItemReject{}, donated_item.ErrorDonatedItemNotFound
	}

	if donatedItemEntity.Status.Status != donated_item.StatusRequested {
		return response_donation.DonationItemReject{}, donated_item.ErrorInvalidStatusTransition
	}

	rejectedDonatedItemEntity := donatedItemEntity
	rejectedDonatedItemEntity.ID = identity.NewID(identity.NilID.String())
	rejectedDonatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusRejected)
	if err != nil {
		return response_donation.DonationItemReject{}, donated_item.ErrorInvalidStatus
	}
	rejectedDonatedItem, err := s.donatedItemRepository.Create(ctx, tx, rejectedDonatedItemEntity)
	if err != nil {
		return response_donation.DonationItemReject{}, donated_item.ErrorRejectDonatedItem
	}

	donatedItemEntity.RecipientID = identity.NewID(identity.NilID.String())
	donatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusOpened)
	if err != nil {
		return response_donation.DonationItemReject{}, donated_item.ErrorInvalidStatus
	}

	_, err = s.donatedItemRepository.Update(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemReject{}, donated_item.ErrorRejectDonatedItem
	}

	return response_donation.DonationItemReject{
		ID:          rejectedDonatedItem.ID.String(),
		DonorID:     rejectedDonatedItem.DonorID.String(),
		RecipientID: rejectedDonatedItem.RecipientID.String(),
		Status:      rejectedDonatedItem.Status.Status,
		Name:        rejectedDonatedItem.Name,
		Description: rejectedDonatedItem.Description,
		Condition:   rejectedDonatedItem.Condition.Value,
		PickCity:    rejectedDonatedItem.PickCity,
		PickAddress: rejectedDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) TakenDonatedItem(ctx context.Context, req request_donation.DonationItemTaken) (response_donation.DonationItemTaken, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonationItemTaken{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonationItemTaken{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	donatedItemEntity, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.ID)
	if err != nil {
		return response_donation.DonationItemTaken{}, donated_item.ErrorDonatedItemNotFound
	}

	if donatedItemEntity.Status.Status != donated_item.StatusAccepted {
		return response_donation.DonationItemTaken{}, donated_item.ErrorInvalidStatusTransition
	}

	donatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusTaken)
	if err != nil {
		return response_donation.DonationItemTaken{}, donated_item.ErrorInvalidStatus
	}

	updatedDonatedItem, err := s.donatedItemRepository.Update(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemTaken{}, donated_item.ErrorTakenDonatedItem
	}

	return response_donation.DonationItemTaken{
		ID:          updatedDonatedItem.ID.String(),
		DonorID:     updatedDonatedItem.DonorID.String(),
		RecipientID: updatedDonatedItem.RecipientID.String(),
		Status:      updatedDonatedItem.Status.Status,
		Name:        updatedDonatedItem.Name,
		Description: updatedDonatedItem.Description,
		Condition:   updatedDonatedItem.Condition.Value,
		PickCity:    updatedDonatedItem.PickCity,
		PickAddress: updatedDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) UpdateDonatedItem(ctx context.Context, req request_donation.DonationItemUpdate) (response_donation.DonationItemUpdate, error) {
	condition, err := shared.NewLikertScale(req.Condition)
	if err != nil {
		return response_donation.DonationItemUpdate{}, donated_item.ErrorUpdateDonatedItem
	}

	donatedItemEntity := donated_item.DonatedItem{
		ID:          identity.NewID(req.ID),
		Name:        req.Name,
		Description: req.Description,
		Condition:   condition,
		PickAddress: req.PickAddress,
	}

	updatedDonatedItem, err := s.donatedItemRepository.Update(ctx, nil, donatedItemEntity)
	if err != nil {
		return response_donation.DonationItemUpdate{}, donated_item.ErrorUpdateDonatedItem
	}

	return response_donation.DonationItemUpdate{
		ID:          updatedDonatedItem.ID.String(),
		Name:        updatedDonatedItem.Name,
		Description: updatedDonatedItem.Description,
		Condition:   updatedDonatedItem.Condition.Value,
		PickAddress: updatedDonatedItem.PickAddress,
	}, nil
}

func (s *donationService) DeleteDonatedItem(ctx context.Context, id string) error {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedDonatedItem, err := s.donatedItemRepository.GetDonatedItemByID(ctx, nil, id)
	if err != nil {
		return donated_item.ErrorDonatedItemNotFound
	}

	err = s.donatedItemRepository.Delete(ctx, tx, retrievedDonatedItem.ID.String())
	if err != nil {
		return donated_item.ErrorDeleteDonatedItem
	}

	return nil
}
