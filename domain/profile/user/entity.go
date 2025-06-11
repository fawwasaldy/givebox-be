package user

import (
	"givebox/domain/identity"
	"givebox/domain/shared"
)

type User struct {
	ID          identity.ID
	Username    string
	Password    Password
	FullName    string
	PhoneNumber string
	shared.Timestamp
}
