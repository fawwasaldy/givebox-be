package service

import (
	"context"
	"errors"
	"fmt"
	"givebox/application/request"
	"givebox/application/response"
	user2 "givebox/domain/profile/user"
	"givebox/domain/refresh_token"
	"givebox/domain/shared"
	"givebox/domain/user"
	"givebox/infrastructure/database/validation"
	"givebox/platform/pagination"
	"gorm.io/gorm"
	"time"
)

type (
	UserService interface {
		Register(ctx context.Context, req request.UserRegister) (response.UserCreate, error)
		GetAllUsersWithPagination(ctx context.Context, req pagination.Request) (pagination.ResponseWithData, error)
		GetUserByID(ctx context.Context, userID string) (response.User, error)
		GetUserByEmail(ctx context.Context, email string) (response.User, error)
		Update(ctx context.Context, userID string, req request.UserUpdate) (response.UserUpdate, error)
		Delete(ctx context.Context, userID string) error
		Verify(ctx context.Context, req request.UserLogin) (response.RefreshToken, error)
		RefreshToken(ctx context.Context, req request.RefreshToken) (response.RefreshToken, error)
		RevokeRefreshToken(ctx context.Context, userID string) error
	}

	userService struct {
		userRepository         user2.Repository
		refreshTokenRepository refresh_token.Repository
		userDomainService      user.Service
		jwtService             JWTService
		transaction            interface{}
	}
)

func NewUserService(
	userRepository user2.Repository,
	refreshTokenRepository refresh_token.Repository,
	userDomainService user.Service,
	jwtService JWTService,
	transaction interface{},
) UserService {
	return &userService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		userDomainService:      userDomainService,
		jwtService:             jwtService,
		transaction:            transaction,
	}
}

func (s *userService) Register(ctx context.Context, req request.UserRegister) (response.UserCreate, error) {
	var filename string

	_, flag, err := s.userRepository.CheckEmail(ctx, nil, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.UserCreate{}, err
	}

	if flag {
		return response.UserCreate{}, user2.ErrorEmailAlreadyExists
	}

	if req.Image != nil {
		filename, err = s.userDomainService.UploadImage(req.Image)
		if err != nil {
			return response.UserCreate{}, err
		}
	}

	password, err := user2.NewPassword(req.Password)
	if err != nil {
		return response.UserCreate{}, err
	}
	role, err := user.NewRole(user.RoleUser)
	if err != nil {
		return response.UserCreate{}, err
	}
	imageUrl, err := shared.NewURL(filename)
	if err != nil {
		return response.UserCreate{}, err
	}

	userEntity := user2.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    password,
		Role:        role,
		ImageUrl:    imageUrl,
		IsVerified:  false,
	}

	registeredUser, err := s.userRepository.Register(ctx, nil, userEntity)
	if err != nil {
		return response.UserCreate{}, user2.ErrorCreateUser
	}

	return response.UserCreate{
		ID:          registeredUser.ID.String(),
		Name:        registeredUser.Name,
		Email:       registeredUser.Email,
		PhoneNumber: registeredUser.PhoneNumber,
		Role:        registeredUser.Role.Name,
		ImageUrl:    registeredUser.ImageUrl.Path,
		IsVerified:  registeredUser.IsVerified,
	}, nil
}

func (s *userService) GetAllUsersWithPagination(ctx context.Context, req pagination.Request) (pagination.ResponseWithData, error) {
	retrievedData, err := s.userRepository.GetAllUsersWithPagination(ctx, nil, req)
	if err != nil {
		return pagination.ResponseWithData{}, user2.ErrorGetAllUsers
	}

	data := make([]any, 0, len(retrievedData.Data))
	for _, retrievedUser := range retrievedData.Data {
		userEntity, ok := retrievedUser.(user2.User)
		if !ok {
			return pagination.ResponseWithData{}, errors.New("failed to cast retrieved data to user.User")
		}
		data = append(data, response.User{
			ID:          userEntity.ID.String(),
			Name:        userEntity.Name,
			Email:       userEntity.Email,
			PhoneNumber: userEntity.PhoneNumber,
			Role:        userEntity.Role.Name,
			ImageUrl:    userEntity.ImageUrl.Path,
			IsVerified:  userEntity.IsVerified,
		})
	}

	retrievedData = pagination.ResponseWithData{
		Data:     data,
		Response: retrievedData.Response,
	}
	return retrievedData, nil
}

func (s *userService) GetUserByID(ctx context.Context, userID string) (response.User, error) {
	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return response.User{}, user2.ErrorGetUserById
	}

	return response.User{
		ID:          retrievedUser.ID.String(),
		Name:        retrievedUser.Name,
		Email:       retrievedUser.Email,
		PhoneNumber: retrievedUser.PhoneNumber,
		Role:        retrievedUser.Role.Name,
		ImageUrl:    retrievedUser.ImageUrl.Path,
		IsVerified:  retrievedUser.IsVerified,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (response.User, error) {
	retrievedUser, err := s.userRepository.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return response.User{}, user2.ErrorGetUserByEmail
	}

	return response.User{
		ID:          retrievedUser.ID.String(),
		Name:        retrievedUser.Name,
		Email:       retrievedUser.Email,
		PhoneNumber: retrievedUser.PhoneNumber,
		Role:        retrievedUser.Role.Name,
		ImageUrl:    retrievedUser.ImageUrl.Path,
		IsVerified:  retrievedUser.IsVerified,
	}, nil
}

func (s *userService) Update(ctx context.Context, userID string, req request.UserUpdate) (response.UserUpdate, error) {
	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return response.UserUpdate{}, user2.ErrorUserNotFound
	}

	userEntity := user2.User{
		ID:          retrievedUser.ID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        retrievedUser.Role,
	}

	updatedUser, err := s.userRepository.Update(ctx, nil, userEntity)
	if err != nil {
		return response.UserUpdate{}, user2.ErrorUpdateUser
	}

	return response.UserUpdate{
		ID:          updatedUser.ID.String(),
		Name:        updatedUser.Name,
		Email:       updatedUser.Email,
		PhoneNumber: updatedUser.PhoneNumber,
		Role:        updatedUser.Role.Name,
		IsVerified:  updatedUser.IsVerified,
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
			err = RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedUser, err := s.userRepository.GetUserByID(ctx, nil, userID)
	if err != nil {
		return user2.ErrorUserNotFound
	}

	err = s.userRepository.Delete(ctx, tx, retrievedUser.ID.String())
	err = fmt.Errorf("test error")
	if err != nil {
		return user2.ErrorDeleteUser
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req request.UserLogin) (response.RefreshToken, error) {
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
			err = RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedUser, err := s.userRepository.GetUserByEmail(ctx, tx, req.Email)
	if err != nil {
		return response.RefreshToken{}, user2.ErrorEmailNotFound
	}

	checkPassword, err := retrievedUser.Password.IsPasswordMatch([]byte(req.Password))
	if err != nil || !checkPassword {
		return response.RefreshToken{}, err
	}

	accessToken := s.jwtService.GenerateAccessToken(retrievedUser.ID.String(), retrievedUser.Role.Name)

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
		Role:         retrievedUser.Role.Name,
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
			err = RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	retrievedRefreshToken, err := s.refreshTokenRepository.FindByToken(ctx, tx, req.RefreshToken)
	if err != nil {
		return response.RefreshToken{}, user2.ErrorTokenInvalid
	}

	if time.Now().After(retrievedRefreshToken.ExpiresAt) {
		return response.RefreshToken{}, user2.ErrorTokenExpired
	}

	retrievedUser, err := s.userRepository.GetUserByID(ctx, tx, retrievedRefreshToken.UserID.String())
	if err != nil {
		return response.RefreshToken{}, user2.ErrorUserNotFound
	}

	accessToken := s.jwtService.GenerateAccessToken(retrievedUser.ID.String(), retrievedUser.Role.Name)

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
		Role:         retrievedUser.Role.Name,
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
			err = RecoveredFromPanic(r)
		}
		validatedTransaction.CommitOrRollback(ctx, tx, err)
	}()

	_, err = s.userRepository.GetUserByID(ctx, tx, userID)
	if err != nil {
		return user2.ErrorUserNotFound
	}

	if err = s.refreshTokenRepository.DeleteByUserID(ctx, tx, userID); err != nil {
		return err
	}

	return nil
}
