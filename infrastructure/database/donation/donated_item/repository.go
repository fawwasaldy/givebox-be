package donated_item

import (
	"context"
	"givebox/domain/donation/donated_item"
	"givebox/infrastructure/database/transaction"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) donated_item.Repository {
	return &repository{db: db}
}

func (r repository) GetAllDonatedItemsWithPagination(ctx context.Context, tx interface{}, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var donatedItemSchemas []DonatedItem
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&DonatedItem{})
	if req.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR pick_city ILIKE ? OR pick_address ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&donatedItemSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	donatedItemEntities := make([]any, len(donatedItemSchemas))
	for i, donatedItemSchema := range donatedItemSchemas {
		donatedItemEntities[i] = SchemaToEntity(donatedItemSchema)
	}
	return pagination.ResponseWithData{
		Data: donatedItemEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) GetAllDonatedItemsByCategoryIDWithPagination(ctx context.Context, tx interface{}, categoryID string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var donatedItemSchemas []DonatedItem
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&DonatedItem{}).
		Joins("JOIN donated_items_categories ON donated_items.id = donated_items_categories.donated_item_id").
		Where("donated_items_categories.category_id = ?", categoryID)
	if req.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR pick_city ILIKE ? OR pick_address ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&donatedItemSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	donatedItemEntities := make([]any, len(donatedItemSchemas))
	for i, donatedItemSchema := range donatedItemSchemas {
		donatedItemEntities[i] = SchemaToEntity(donatedItemSchema)
	}
	return pagination.ResponseWithData{
		Data: donatedItemEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) GetAllDonatedItemsByCityWithPagination(ctx context.Context, tx interface{}, city string, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var donatedItemSchemas []DonatedItem
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&DonatedItem{}).Where("pick_city ILIKE ?", "%"+city+"%")
	if req.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR pick_address ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&donatedItemSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	donatedItemEntities := make([]any, len(donatedItemSchemas))
	for i, donatedItemSchema := range donatedItemSchemas {
		donatedItemEntities[i] = SchemaToEntity(donatedItemSchema)
	}
	return pagination.ResponseWithData{
		Data: donatedItemEntities,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, nil
}

func (r repository) GetDonatedItemByID(ctx context.Context, tx interface{}, id string) (donated_item.DonatedItem, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item.DonatedItem{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var donatedItemSchema DonatedItem
	if err = db.WithContext(ctx).Where("id = ?", id).Take(&donatedItemSchema).Error; err != nil {
		return donated_item.DonatedItem{}, err
	}

	donatedItemEntity := SchemaToEntity(donatedItemSchema)
	return donatedItemEntity, nil
}

func (r repository) CountDonatedItemsByCategoryID(ctx context.Context, tx interface{}, categoryID string) (int64, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return 0, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var count int64
	if err = db.WithContext(ctx).Model(&DonatedItem{}).
		Where("category_id = ?", categoryID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r repository) Create(ctx context.Context, tx interface{}, donatedItemEntity donated_item.DonatedItem) (donated_item.DonatedItem, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item.DonatedItem{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	donatedItemSchema := EntityToSchema(donatedItemEntity)
	if err = db.WithContext(ctx).Create(&donatedItemSchema).Error; err != nil {
		return donated_item.DonatedItem{}, err
	}

	donatedItemEntity = SchemaToEntity(donatedItemSchema)
	return donatedItemEntity, nil
}

func (r repository) Update(ctx context.Context, tx interface{}, donatedItemEntity donated_item.DonatedItem) (donated_item.DonatedItem, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return donated_item.DonatedItem{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	donatedItemSchema := EntityToSchema(donatedItemEntity)
	if err = db.WithContext(ctx).Updates(&donatedItemSchema).Error; err != nil {
		return donated_item.DonatedItem{}, err
	}

	donatedItemEntity = SchemaToEntity(donatedItemSchema)
	return donatedItemEntity, nil
}
