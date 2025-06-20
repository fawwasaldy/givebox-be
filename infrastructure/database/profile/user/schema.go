package user

import (
	"github.com/google/uuid"
	"givebox/domain/identity"
	"givebox/domain/profile/user"
	"givebox/domain/shared"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	Biography   string    `gorm:"type:text;column:biography"`
	FirstName   string    `gorm:"type:varchar(50);not null;column:first_name"`
	LastName    string    `gorm:"type:varchar(50);not null;column:last_name"`
	Email       string    `gorm:"type:varchar(50);uniqueIndex;not null;column:email"`
	Password    string    `gorm:"type:varchar(255);not null;column:password"`
	PhoneNumber string    `gorm:"type:varchar(20);index;column:phone_number"`
	City        string    `gorm:"type:varchar(100);column:city"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "users"
}

func EntityToSchema(entity user.User) User {
	return User{
		ID:          entity.ID.ID,
		Biography:   entity.Biography,
		FirstName:   entity.Name.FirstName,
		LastName:    entity.Name.LastName,
		Email:       entity.Email,
		Password:    entity.Password.Password,
		PhoneNumber: entity.PhoneNumber,
		City:        entity.City,
		CreatedAt:   entity.Timestamp.CreatedAt,
		UpdatedAt:   entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema User) user.User {
	return user.User{
		ID:          identity.NewIDFromSchema(schema.ID),
		Biography:   schema.Biography,
		Name:        user.NewNameFromSchema(schema.FirstName, schema.LastName),
		Email:       schema.Email,
		Password:    user.NewPasswordFromSchema(schema.Password),
		PhoneNumber: schema.PhoneNumber,
		City:        schema.City,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
