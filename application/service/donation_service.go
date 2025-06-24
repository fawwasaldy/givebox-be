package service

import (
	"context"
	"givebox/application"
	request_donation "givebox/application/request/donation"
	response_donation "givebox/application/response/donation"
	"givebox/domain/donation/category"
	"givebox/domain/donation/donated_item"
	"givebox/domain/donation/donated_item_recipient"
	"givebox/domain/donation/image"
	"givebox/domain/identity"
	"givebox/domain/profile/user"
	"givebox/domain/shared"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type (
	DonationService interface {
		GetAllDonatedItemsWithPagination(ctx context.Context, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByCategoryIDWithPagination(ctx context.Context, categoryID string, req pagination.Request) (pagination.ResponseWithData, error)
		GetAllDonatedItemsByCityWithPagination(ctx context.Context, city string, req pagination.Request) (pagination.ResponseWithData, error)
		GetDonatedItemByID(ctx context.Context, id string) (response_donation.DetailedDonatedItem, error)
		OpenDonatedItem(ctx context.Context, donorID string, req request_donation.DonatedItemOpen) (response_donation.DonatedItemOpen, error)
		AcceptDonatedItem(ctx context.Context, req request_donation.DonatedItemAccept) (response_donation.DonatedItemAccept, error)
		GetAllImagesByDonatedItemID(ctx context.Context, donatedItemID string) ([]response_donation.Image, error)
		GetAllCategories(ctx context.Context) ([]response_donation.Category, error)
	}

	donationService struct {
		donatedItemRepository          donated_item.Repository
		donatedItemRecipientRepository donated_item_recipient.Repository
		imageRepository                image.Repository
		categoryRepository             category.Repository
		userRepository                 user.Repository
		transaction                    interface{}
	}
)

func NewDonationService(
	donatedItemRepository donated_item.Repository,
	donatedItemRecipientRepository donated_item_recipient.Repository,
	imageRepository image.Repository,
	categoryRepository category.Repository,
	userRepository user.Repository,
	transaction interface{},
) DonationService {
	return &donationService{
		donatedItemRepository:          donatedItemRepository,
		donatedItemRecipientRepository: donatedItemRecipientRepository,
		imageRepository:                imageRepository,
		categoryRepository:             categoryRepository,
		userRepository:                 userRepository,
		transaction:                    transaction,
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

		var retrievedDonor user.User
		retrievedDonor, err = s.userRepository.GetUserByID(ctx, nil, donatedItemEntity.DonorID.String())
		if err != nil {
			return pagination.ResponseWithData{}, user.ErrorGetUserById
		}
		var retrievedCategory category.Category
		retrievedCategory, err = s.categoryRepository.GetCategoryByID(ctx, nil, donatedItemEntity.CategoryID.String())
		if err != nil {
			return pagination.ResponseWithData{}, category.ErrorGetCategoryById
		}

		data = append(data, response_donation.DonatedItem{
			ID:          donatedItemEntity.ID.String(),
			DonorName:   retrievedDonor.Name.FullName(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Category:    retrievedCategory.Name,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
			IsUrgent:    donatedItemEntity.IsUrgent,
			CreatedAt:   donatedItemEntity.CreatedAt.String(),
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

		var retrievedDonor user.User
		retrievedDonor, err = s.userRepository.GetUserByID(ctx, nil, donatedItemEntity.DonorID.String())
		if err != nil {
			return pagination.ResponseWithData{}, user.ErrorGetUserById
		}
		var retrievedCategory category.Category
		retrievedCategory, err = s.categoryRepository.GetCategoryByID(ctx, nil, donatedItemEntity.CategoryID.String())
		if err != nil {
			return pagination.ResponseWithData{}, category.ErrorGetCategoryById
		}

		data = append(data, response_donation.DonatedItem{
			ID:          donatedItemEntity.ID.String(),
			DonorName:   retrievedDonor.Name.FullName(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Category:    retrievedCategory.Name,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
			IsUrgent:    donatedItemEntity.IsUrgent,
			CreatedAt:   donatedItemEntity.CreatedAt.String(),
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetAllDonatedItemsByCityWithPagination(ctx context.Context, city string, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.donatedItemRepository.GetAllDonatedItemsByCityWithPagination(ctx, nil, city, req)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedDonatedItems := range retrievedData.Data {
		donatedItemEntity, ok := retrievedDonatedItems.(donated_item.DonatedItem)
		if !ok {
			return pagination.ResponseWithData{}, donated_item.ErrorInvalidDonatedItemType
		}

		var retrievedDonor user.User
		retrievedDonor, err = s.userRepository.GetUserByID(ctx, nil, donatedItemEntity.DonorID.String())
		if err != nil {
			return pagination.ResponseWithData{}, user.ErrorGetUserById
		}
		var retrievedCategory category.Category
		retrievedCategory, err = s.categoryRepository.GetCategoryByID(ctx, nil, donatedItemEntity.CategoryID.String())
		if err != nil {
			return pagination.ResponseWithData{}, category.ErrorGetCategoryById
		}

		data = append(data, response_donation.DonatedItem{
			ID:          donatedItemEntity.ID.String(),
			DonorName:   retrievedDonor.Name.FullName(),
			Name:        donatedItemEntity.Name,
			Description: donatedItemEntity.Description,
			Category:    retrievedCategory.Name,
			Condition:   donatedItemEntity.Condition.Value,
			PickCity:    donatedItemEntity.PickCity,
			IsUrgent:    donatedItemEntity.IsUrgent,
			CreatedAt:   donatedItemEntity.CreatedAt.String(),
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}

	return retrievedData, nil
}

func (s *donationService) GetDonatedItemByID(ctx context.Context, id string) (response_donation.DetailedDonatedItem, error) {
	retrievedDonatedItem, err := s.donatedItemRepository.GetDonatedItemByID(ctx, nil, id)
	if err != nil {
		return response_donation.DetailedDonatedItem{}, donated_item.ErrorGetDonatedItemById
	}

	retrievedDonor, err := s.userRepository.GetUserByID(ctx, nil, retrievedDonatedItem.DonorID.String())
	if err != nil {
		return response_donation.DetailedDonatedItem{}, user.ErrorGetUserById
	}
	retrievedCategory, err := s.categoryRepository.GetCategoryByID(ctx, nil, retrievedDonatedItem.CategoryID.String())
	if err != nil {
		return response_donation.DetailedDonatedItem{}, category.ErrorGetCategoryById
	}

	return response_donation.DetailedDonatedItem{
		ID:                  retrievedDonatedItem.ID.String(),
		DonorName:           retrievedDonor.Name.FullName(),
		Name:                retrievedDonatedItem.Name,
		Description:         retrievedDonatedItem.Description,
		Category:            retrievedCategory.Name,
		Condition:           retrievedDonatedItem.Condition.Value,
		QuantityDescription: retrievedDonatedItem.QuantityDescription,
		PickCity:            retrievedDonatedItem.PickCity,
		PickAddress:         retrievedDonatedItem.PickAddress,
		PickingStatus:       retrievedDonatedItem.PickingStatus.Status,
		DeliveryTime:        retrievedDonatedItem.DeliveryTime,
		IsUrgent:            retrievedDonatedItem.IsUrgent,
		AdditionalNote:      retrievedDonatedItem.AdditionalNote,
	}, nil
}

func (s *donationService) OpenDonatedItem(ctx context.Context, donorID string, req request_donation.DonatedItemOpen) (response_donation.DonatedItemOpen, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonatedItemOpen{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonatedItemOpen{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	condition, err := shared.NewLikertScale(req.Condition)
	if err != nil {
		return response_donation.DonatedItemOpen{}, donated_item.ErrorOpenDonatedItem
	}
	status, err := donated_item.NewStatus(donated_item.StatusOpened)
	if err != nil {
		return response_donation.DonatedItemOpen{}, donated_item.ErrorInvalidStatus
	}
	pickingStatus, err := donated_item.NewPickingStatus(req.PickingStatus)
	if err != nil {
		return response_donation.DonatedItemOpen{}, donated_item.ErrorInvalidPickingStatus
	}

	donatedItemEntity := donated_item.DonatedItem{
		DonorID:             identity.NewID(donorID),
		CategoryID:          identity.NewID(req.CategoryID),
		Status:              status,
		Name:                req.Name,
		Description:         req.Description,
		Condition:           condition,
		QuantityDescription: req.QuantityDescription,
		PickCity:            req.PickCity,
		PickAddress:         req.PickAddress,
		PickingStatus:       pickingStatus,
		DeliveryTime:        req.DeliveryTime,
		IsUrgent:            req.IsUrgent,
		AdditionalNote:      req.AdditionalNote,
	}

	createdDonatedItem, err := s.donatedItemRepository.Create(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonatedItemOpen{}, donated_item.ErrorOpenDonatedItem
	}

	for _, reqImage := range req.Images {
		var imageURL shared.URL
		imageURL, err = shared.NewURL(reqImage)
		if err != nil {
			return response_donation.DonatedItemOpen{}, donated_item.ErrorOpenDonatedItem
		}
		imageEntity := image.Image{
			DonatedItemID: createdDonatedItem.ID,
			ImageURL:      imageURL,
		}

		if _, err = s.imageRepository.Create(ctx, tx, imageEntity); err != nil {
			return response_donation.DonatedItemOpen{}, donated_item.ErrorOpenDonatedItem
		}
	}

	retrievedDonor, err := s.userRepository.GetUserByID(ctx, tx, createdDonatedItem.DonorID.String())
	if err != nil {
		return response_donation.DonatedItemOpen{}, user.ErrorGetUserById
	}
	retrievedCategory, err := s.categoryRepository.GetCategoryByID(ctx, tx, createdDonatedItem.CategoryID.String())
	if err != nil {
		return response_donation.DonatedItemOpen{}, category.ErrorGetCategoryById
	}

	return response_donation.DonatedItemOpen{
		ID:                  createdDonatedItem.ID.String(),
		DonorName:           retrievedDonor.Name.FullName(),
		Status:              createdDonatedItem.Status.Status,
		Name:                createdDonatedItem.Name,
		Description:         createdDonatedItem.Description,
		Category:            retrievedCategory.Name,
		Condition:           createdDonatedItem.Condition.Value,
		QuantityDescription: createdDonatedItem.QuantityDescription,
		PickCity:            createdDonatedItem.PickCity,
		PickAddress:         createdDonatedItem.PickAddress,
		PickingStatus:       createdDonatedItem.PickingStatus.Status,
		DeliveryTime:        createdDonatedItem.DeliveryTime,
		IsUrgent:            createdDonatedItem.IsUrgent,
		AdditionalNote:      createdDonatedItem.AdditionalNote,
	}, nil
}

func (s *donationService) AcceptDonatedItem(ctx context.Context, req request_donation.DonatedItemAccept) (response_donation.DonatedItemAccept, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_donation.DonatedItemAccept{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_donation.DonatedItemAccept{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	donatedItemEntity, err := s.donatedItemRepository.GetDonatedItemByID(ctx, tx, req.ID)
	if err != nil {
		return response_donation.DonatedItemAccept{}, donated_item.ErrorDonatedItemNotFound
	}

	if donatedItemEntity.Status.Status != donated_item.StatusOpened {
		return response_donation.DonatedItemAccept{}, donated_item.ErrorInvalidStatusTransition
	}

	donatedItemEntity.Status, err = donated_item.NewStatus(donated_item.StatusAccepted)
	if err != nil {
		return response_donation.DonatedItemAccept{}, donated_item.ErrorInvalidStatus
	}

	updatedDonatedItem, err := s.donatedItemRepository.Update(ctx, tx, donatedItemEntity)
	if err != nil {
		return response_donation.DonatedItemAccept{}, donated_item.ErrorAcceptDonatedItem
	}

	donatedItemRecipientEntity := donated_item_recipient.DonatedItemRecipient{
		DonatedItemID: identity.NewID(req.ID),
		RecipientID:   identity.NewID(req.RecipientID),
		IsAccepted:    true,
	}

	_, err = s.donatedItemRecipientRepository.Update(ctx, tx, donatedItemRecipientEntity)
	if err != nil {
		return response_donation.DonatedItemAccept{}, donated_item_recipient.ErrorUpdateDonatedItemRecipient
	}

	retrievedDonor, err := s.userRepository.GetUserByID(ctx, tx, updatedDonatedItem.DonorID.String())
	if err != nil {
		return response_donation.DonatedItemAccept{}, user.ErrorGetUserById
	}
	retrievedRecipient, err := s.userRepository.GetUserByID(ctx, tx, req.RecipientID)
	if err != nil {
		return response_donation.DonatedItemAccept{}, user.ErrorGetUserById
	}
	retrievedCategory, err := s.categoryRepository.GetCategoryByID(ctx, tx, updatedDonatedItem.CategoryID.String())
	if err != nil {
		return response_donation.DonatedItemAccept{}, category.ErrorGetCategoryById
	}

	return response_donation.DonatedItemAccept{
		ID:                  updatedDonatedItem.ID.String(),
		DonorName:           retrievedDonor.Name.FullName(),
		RecipientName:       retrievedRecipient.Name.FullName(),
		Status:              updatedDonatedItem.Status.Status,
		Name:                updatedDonatedItem.Name,
		Description:         updatedDonatedItem.Description,
		Category:            retrievedCategory.Name,
		Condition:           updatedDonatedItem.Condition.Value,
		QuantityDescription: updatedDonatedItem.QuantityDescription,
		PickCity:            updatedDonatedItem.PickCity,
		PickAddress:         updatedDonatedItem.PickAddress,
		PickingStatus:       updatedDonatedItem.PickingStatus.Status,
		DeliveryTime:        updatedDonatedItem.DeliveryTime,
		IsUrgent:            updatedDonatedItem.IsUrgent,
		AdditionalNote:      updatedDonatedItem.AdditionalNote,
	}, nil
}

func (s *donationService) GetAllImagesByDonatedItemID(ctx context.Context, donatedItemID string) ([]response_donation.Image, error) {
	retrievedImages, err := s.imageRepository.GetAllImagesByDonatedItemID(ctx, nil, donatedItemID)
	if err != nil {
		return nil, image.ErrorGetAllImages
	}

	data := make([]response_donation.Image, 0, len(retrievedImages))
	for _, retrievedImage := range retrievedImages {
		data = append(data, response_donation.Image{
			ID:            retrievedImage.ID.String(),
			DonatedItemID: retrievedImage.DonatedItemID.String(),
			ImageURL:      retrievedImage.ImageURL.Path,
		})
	}

	return data, nil
}

func (s *donationService) GetAllCategories(ctx context.Context) ([]response_donation.Category, error) {
	retrievedCategories, err := s.categoryRepository.GetAllCategories(ctx, nil)
	if err != nil {
		return nil, category.ErrorGetAllCategories
	}

	data := make([]response_donation.Category, 0, len(retrievedCategories))
	for _, retrievedCategory := range retrievedCategories {
		data = append(data, response_donation.Category{
			ID:   retrievedCategory.ID.String(),
			Name: retrievedCategory.Name,
		})
	}

	return data, nil
}
