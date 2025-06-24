package service

import (
	"context"
	response_donation "givebox/application/response/donation"
	"givebox/domain/donation/donated_item"
	"givebox/domain/profile/user"
	infrastructure_donated_item "givebox/infrastructure/database/donation/donated_item"
	infrastructure_user "givebox/infrastructure/database/profile/user"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
)

type (
	AnalyticService interface {
		GetTotalDonatedItems(ctx context.Context) (int64, error)
		GetTotalUsers(ctx context.Context) (int64, error)
		GetTotalOpenedDonatedItems(ctx context.Context) (int64, error)
		GetAcceptedDonatedItemPercentage(ctx context.Context) (float64, error)
		GetSixCategoriesByMostDonatedItems(ctx context.Context) ([]response_donation.CategoryAnalytic, error)
		GetThreeLatestDonatedItems(ctx context.Context) ([]response_donation.DonatedItem, error)
	}

	analyticService struct {
		donatedItemRepository donated_item.Repository
		userRepository        user.Repository
		transaction           *transaction.Repository
	}
)

func NewAnalyticService(
	donatedItemRepository donated_item.Repository,
	userRepository user.Repository,
	transaction *transaction.Repository,
) AnalyticService {
	return &analyticService{
		donatedItemRepository: donatedItemRepository,
		userRepository:        userRepository,
		transaction:           transaction,
	}
}

func (s *analyticService) GetTotalDonatedItems(ctx context.Context) (int64, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return 0, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var count int64
	if err = db.WithContext(ctx).Model(&infrastructure_donated_item.DonatedItem{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *analyticService) GetTotalUsers(ctx context.Context) (int64, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return 0, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var count int64
	if err = db.WithContext(ctx).Model(&infrastructure_user.User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *analyticService) GetTotalOpenedDonatedItems(ctx context.Context) (int64, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return 0, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var count int64
	if err = db.WithContext(ctx).Model(&infrastructure_donated_item.DonatedItem{}).
		Where("status = ?", donated_item.StatusOpened).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *analyticService) GetAcceptedDonatedItemPercentage(ctx context.Context) (float64, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return 0, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var totalAccepted int64
	if err = db.WithContext(ctx).Model(&infrastructure_donated_item.DonatedItem{}).
		Where("status = ?", donated_item.StatusAccepted).
		Count(&totalAccepted).Error; err != nil {
		return 0, err
	}

	var totalDonatedItems int64
	if err = db.WithContext(ctx).Model(&infrastructure_donated_item.DonatedItem{}).Count(&totalDonatedItems).Error; err != nil {
		return 0, err
	}

	if totalDonatedItems == 0 {
		return 0, nil
	}

	return float64(totalAccepted) / float64(totalDonatedItems) * 100, nil
}

func (s *analyticService) GetSixCategoriesByMostDonatedItems(ctx context.Context) ([]response_donation.CategoryAnalytic, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return nil, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var categories []response_donation.CategoryAnalytic
	if err = db.WithContext(ctx).
		Model(&infrastructure_donated_item.DonatedItem{}).
		Joins("RIGHT JOIN categories ON categories.id = donated_items.category_id").
		Group("categories.id").
		Order("COUNT(donated_items.id) DESC").
		Limit(6).
		Select("categories.id, categories.name, COUNT(donated_items.id) AS count").
		Scan(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *analyticService) GetThreeLatestDonatedItems(ctx context.Context) ([]response_donation.DonatedItem, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return nil, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = s.transaction.DB()
	}

	var donatedItems []response_donation.DonatedItem
	var queryResult []struct {
		ID             string
		DonorFirstName string
		DonorLastName  string
		Name           string
		Description    string
		CategoryName   string
		Condition      int
		PickCity       string
		IsUrgent       bool
		CreatedAt      string
	}
	if err = db.WithContext(ctx).
		Model(&infrastructure_donated_item.DonatedItem{}).
		Joins("JOIN users ON users.id = donated_items.donor_id").
		Joins("JOIN categories ON categories.id = donated_items.category_id").
		Order("created_at DESC").
		Limit(3).
		Select("donated_items.id, users.first_name AS donor_first_name, users.last_name AS donor_last_name, donated_items.name, donated_items.description, categories.name AS category_name, donated_items.condition, donated_items.pick_city, donated_items.is_urgent, donated_items.created_at").
		Scan(&queryResult).Error; err != nil {
		return nil, err
	}

	for _, item := range queryResult {
		donatedItem := response_donation.DonatedItem{
			ID:          item.ID,
			DonorName:   item.DonorFirstName + " " + item.DonorLastName,
			Name:        item.Name,
			Description: item.Description,
			Category:    item.CategoryName,
			Condition:   item.Condition,
			PickCity:    item.PickCity,
			IsUrgent:    item.IsUrgent,
			CreatedAt:   item.CreatedAt,
		}
		donatedItems = append(donatedItems, donatedItem)
	}

	return donatedItems, nil
}
