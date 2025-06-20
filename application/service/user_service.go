package service

import (
	"context"
	"errors"
	"givebox/application"
	"givebox/application/request"
	request_profile "givebox/application/request/profile"
	"givebox/application/response"
	response_profile "givebox/application/response/profile"
	"givebox/domain/profile/user"
	"givebox/domain/refresh_token"
	"givebox/infrastructure/database/validation"
	"gorm.io/gorm"
	"time"
)

type (
	UserService interface {
		Register(ctx context.Context, req request_profile.UserRegister) (response_profile.UserCreate, error)
		GetUserByID(ctx context.Context, userID string) (response_profile.User, error)
		GetUserByEmail(ctx context.Context, email string) (response_profile.User, error)
		Update(ctx context.Context, userID string, req request_profile.UserUpdate) (response_profile.UserUpdate, error)
		Delete(ctx context.Context, userID string) error
		Verify(ctx context.Context, req request_profile.UserLogin) (response.RefreshToken, error)
		ChangePassword(ctx context.Context, userID string, req request_profile.UserChangePassword) (response_profile.UserChangePassword, error)
		RefreshToken(ctx context.Context, req request.RefreshToken) (response.RefreshToken, error)
		RevokeRefreshToken(ctx context.Context, userID string) error
	}

	userService struct {
		userRepository         user.Repository
		refreshTokenRepository refresh_token.Repository
		jwtService             JWTService
		transaction            interface{}
	}
)

func NewUserService(
	userRepository user.Repository,
	refreshTokenRepository refresh_token.Repository,
	jwtService JWTService,
	transaction interface{},
) UserService {
	return &userService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		jwtService:             jwtService,
		transaction:            transaction,
	}
}

func (s *userService) Register(ctx context.Context, req request_profile.UserRegister) (response_profile.UserCreate, error) {
	_, flag, err := s.userRepository.CheckEmail(ctx, nil, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response_profile.UserCreate{}, err
	}

	if flag {
		return response_profile.UserCreate{}, user.ErrorEmailAlreadyExists
	}

	name, err := user.NewName(req.FirstName, req.LastName)
	if err != nil {
		return response_profile.UserCreate{}, err
	}
	password, err := user.NewPassword(req.Password)
	if err != nil {
		return response_profile.UserCreate{}, err
	}

	userEntity := user.User{
		Email:       req.Email,
		Name:        name,
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		Password:    password,
	}

	registeredUser, err := s.userRepository.Register(ctx, nil, userEntity)
	if err != nil {
		return response_profile.UserCreate{}, user.ErrorCreateUser
	}

	return response_profile.UserCreate{
		ID:       registeredUser.ID.String(),
		FullName: registeredUser.Name.FullName(),
		Email:    registeredUser.Email,
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, userID string) (response_profile.User, error) {
	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return response_profile.User{}, user.ErrorGetUserById
	}

	return response_profile.User{
		ID:          retrievedUser.ID.String(),
		FullName:    retrievedUser.Name.FullName(),
		Email:       retrievedUser.Email,
		PhoneNumber: retrievedUser.PhoneNumber,
		City:        retrievedUser.City,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (response_profile.User, error) {
	retrievedUser, err := s.userRepository.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return response_profile.User{}, user.ErrorGetUserByEmail
	}

	return response_profile.User{
		ID:          retrievedUser.ID.String(),
		FullName:    retrievedUser.Name.FullName(),
		Email:       retrievedUser.Email,
		PhoneNumber: retrievedUser.PhoneNumber,
		City:        retrievedUser.City,
	}, nil
}

func (s *userService) Update(ctx context.Context, userID string, req request_profile.UserUpdate) (response_profile.UserUpdate, error) {
	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return response_profile.UserUpdate{}, user.ErrorUserNotFound
	}

	var name user.Name
	if req.FirstName != "" && req.LastName != "" {
		name, err = user.NewName(req.FirstName, req.LastName)
		if err != nil {
			return response_profile.UserUpdate{}, err
		}
	}

	userEntity := user.User{
		ID:          retrievedUser.ID,
		Name:        name,
		Biography:   req.Biography,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
	}

	updatedUser, err := s.userRepository.Update(ctx, nil, userEntity)
	if err != nil {
		return response_profile.UserUpdate{}, user.ErrorUpdateUser
	}

	return response_profile.UserUpdate{
		ID:          updatedUser.ID.String(),
		FirstName:   updatedUser.Name.FirstName,
		LastName:    updatedUser.Name.LastName,
		Biography:   updatedUser.Biography,
		Email:       updatedUser.Email,
		PhoneNumber: updatedUser.PhoneNumber,
		City:        updatedUser.City,
	}, nil
}

func (s *userService) Delete(ctx context.Context, userID string) error {
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

	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return user.ErrorUserNotFound
	}

	if err = s.refreshTokenRepository.DeleteByUserID(ctx, tx, userID); err != nil {
		return err
	}

	err = s.userRepository.Delete(ctx, tx, retrievedUser.ID.String())
	if err != nil {
		return user.ErrorDeleteUser
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req request_profile.UserLogin) (response.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response.RefreshToken{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response.RefreshToken{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedUser, err := s.userRepository.GetUserByEmail(ctx, tx, req.Email)
	if err != nil {
		return response.RefreshToken{}, user.ErrorEmailNotFound
	}

	checkPassword, err := retrievedUser.Password.IsPasswordMatch([]byte(req.Password))
	if err != nil || !checkPassword {
		return response.RefreshToken{}, err
	}

	accessToken := s.jwtService.GenerateAccessToken(retrievedUser.ID.String())

	refreshTokenString, expiresAt := s.jwtService.GenerateRefreshToken()

	hashedToken, err := refresh_token.HashToken(refreshTokenString)
	if err != nil {
		return response.RefreshToken{}, err
	}

	if err = s.refreshTokenRepository.DeleteByUserID(ctx, tx, retrievedUser.ID.String()); err != nil {
		return response.RefreshToken{}, err
	}

	refreshTokenEntity := refresh_token.RefreshToken{
		UserID:    retrievedUser.ID,
		Token:     hashedToken,
		ExpiresAt: expiresAt,
	}

	if _, err = s.refreshTokenRepository.Create(ctx, tx, refreshTokenEntity); err != nil {
		return response.RefreshToken{}, err
	}

	return response.RefreshToken{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *userService) ChangePassword(ctx context.Context, userID string, req request_profile.UserChangePassword) (response_profile.UserChangePassword, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response_profile.UserChangePassword{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response_profile.UserChangePassword{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedUser, err := s.userRepository.GetUserByID(ctx, tx, userID)
	if err != nil {
		return response_profile.UserChangePassword{}, user.ErrorUserNotFound
	}

	checkOldPassword, err := retrievedUser.Password.IsPasswordMatch([]byte(req.OldPassword))
	if err != nil || !checkOldPassword {
		return response_profile.UserChangePassword{}, err
	}

	newPassword, err := user.NewPassword(req.NewPassword)
	if err != nil {
		return response_profile.UserChangePassword{}, err
	}

	userEntity := user.User{
		ID:       retrievedUser.ID,
		Password: newPassword,
	}

	updatedUser, err := s.userRepository.Update(ctx, tx, userEntity)
	if err != nil {
		return response_profile.UserChangePassword{}, user.ErrorUpdateUser
	}

	return response_profile.UserChangePassword{
		ID: updatedUser.ID.String(),
	}, nil
}

func (s *userService) RefreshToken(ctx context.Context, req request.RefreshToken) (response.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(s.transaction)
	if err != nil {
		return response.RefreshToken{}, err
	}

	tx, err := validatedTransaction.Begin(ctx)
	if err != nil {
		return response.RefreshToken{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = application.RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedRefreshToken, err := s.refreshTokenRepository.FindByUserID(ctx, tx, req.UserID)
	if err != nil {
		return response.RefreshToken{}, user.ErrorUserNotFound
	}

	if !refresh_token.IsRefreshTokenMatch(req.RefreshToken, retrievedRefreshToken.Token) {
		return response.RefreshToken{}, user.ErrorTokenInvalid
	}

	if time.Now().After(retrievedRefreshToken.ExpiresAt) {
		return response.RefreshToken{}, user.ErrorTokenExpired
	}

	retrievedUser, err := s.userRepository.GetUserByID(ctx, tx, retrievedRefreshToken.UserID.String())
	if err != nil {
		return response.RefreshToken{}, user.ErrorUserNotFound
	}

	accessToken := s.jwtService.GenerateAccessToken(retrievedUser.ID.String())

	refreshTokenString, expiresAt := s.jwtService.GenerateRefreshToken()

	hashedToken, err := refresh_token.HashToken(refreshTokenString)
	if err != nil {
		return response.RefreshToken{}, err
	}

	if err = s.refreshTokenRepository.DeleteByUserID(ctx, tx, retrievedUser.ID.String()); err != nil {
		return response.RefreshToken{}, err
	}

	refreshTokenEntity := refresh_token.RefreshToken{
		UserID:    retrievedUser.ID,
		Token:     hashedToken,
		ExpiresAt: expiresAt,
	}

	if _, err = s.refreshTokenRepository.Create(ctx, tx, refreshTokenEntity); err != nil {
		return response.RefreshToken{}, err
	}

	return response.RefreshToken{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *userService) RevokeRefreshToken(ctx context.Context, userID string) error {
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

	_, err = s.userRepository.GetUserByID(ctx, tx, userID)
	if err != nil {
		return user.ErrorUserNotFound
	}

	if err = s.refreshTokenRepository.DeleteByUserID(ctx, tx, userID); err != nil {
		return err
	}

	return nil
}
