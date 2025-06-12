package user

import "errors"

var (
	ErrorCreateUser            = errors.New("failed to create user")
	ErrorGetUserById           = errors.New("failed to get user by id")
	ErrorGetUserByUsername     = errors.New("failed to get user by username")
	ErrorUsernameAlreadyExists = errors.New("username already exist")
	ErrorUpdateUser            = errors.New("failed to update user")
	ErrorUserNotFound          = errors.New("user not found")
	ErrorUsernameNotFound      = errors.New("username not found")
	ErrorDeleteUser            = errors.New("failed to delete user")
	ErrorTokenInvalid          = errors.New("token invalid")
	ErrorTokenExpired          = errors.New("token expired")
)
