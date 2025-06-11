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
	Username    string    `gorm:"type:varchar(50);uniqueIndex;not null;column:username"`
	Password    string    `gorm:"type:varchar(255);not null;column:password"`
	FullName    string    `gorm:"type:varchar(100);not null;column:full_name"`
	PhoneNumber string    `gorm:"type:varchar(20);index;column:phone_number"`
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
		Username:    entity.Username,
		Password:    entity.Password.Password,
		FullName:    entity.FullName,
		PhoneNumber: entity.PhoneNumber,
		CreatedAt:   entity.Timestamp.CreatedAt,
		UpdatedAt:   entity.Timestamp.UpdatedAt,
	}
}

func SchemaToEntity(schema User) user.User {
	return user.User{
		ID:          identity.NewIDFromSchema(schema.ID),
		Username:    schema.Username,
		Password:    user.NewPasswordFromSchema(schema.Password),
		FullName:    schema.FullName,
		PhoneNumber: schema.PhoneNumber,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
