package service

import (
	"context"
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/model/domain"
	"gallery_go/model/web"
	"gallery_go/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) SignUp(ctx context.Context, request web.UserSignUpRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		FullName: request.FullName,
		Password: hashedPassword,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		result, err := service.UserRepository.FindByUsername(ctx, tx, user.Username)
		if err == nil && result.ID != 0 {
			panic(exception.NewConflictError("Username is already exists"))
		}

		user, err = service.UserRepository.Create(ctx, tx, user)
		helper.PanicIfError(err)

		return nil
	})
	helper.PanicIfError(err)

	jwtToken, err := helper.CreateToken(user)
	helper.PanicIfError(err)

	return web.AuthResponse{
		ID:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Role:      user.Role,
		Token:     jwtToken,
		CreatedAt: user.CreatedAt,
	}
}

func (service *UserServiceImpl) SignIn(ctx context.Context, request web.UserSignInRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindByUsername(ctx, service.DB, request.Username)
	if err != nil {
		panic(exception.NewUnauthorizedError("Invalid credentials"))
	}

	err = helper.ComparePassword(user.Password, request.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError("Invalid credentials"))
	}

	jwtToken, err := helper.CreateToken(user)
	if err != nil {
		panic(err)
	}

	return web.AuthResponse{
		ID:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Role:      user.Role,
		Token:     jwtToken,
		CreatedAt: user.CreatedAt,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := domain.User{
		ID:       request.ID,
		Username: request.Username,
		FullName: request.FullName,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		result, err := service.UserRepository.FindById(ctx, tx, user.ID)
		if err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		result, err = service.UserRepository.FindByUsername(ctx, tx, user.Username)
		if err == nil && result.ID != 0 && result.ID != user.ID {
			panic(exception.NewConflictError("Username is already exists"))
		}

		user, err = service.UserRepository.Update(ctx, tx, user)
		helper.PanicIfError(err)

		return nil
	})
	helper.PanicIfError(err)

	return web.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}
}
