package user

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type User struct {
	ID          identity.ID
	Biography   string
	Name        Name
	Email       string
	Password    Password
	PhoneNumber string
	City        string
	shared.Timestamp
}
