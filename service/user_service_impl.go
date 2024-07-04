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
		ID: user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Role: user.Role,
		Token: jwtToken,
	}
}

func (service *UserServiceImpl) SignIn(ctx context.Context, request web.UserSignInRequest) web.AuthResponse {
	panic("err")
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	panic("err")
}
