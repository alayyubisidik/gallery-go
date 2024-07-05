package test

import (
	"context"
	"gallery_go/app"
	"gallery_go/controller"
	"gallery_go/helper"
	"gallery_go/model/domain"
	"gallery_go/repository"
	"gallery_go/service"
	"net/http"

	"github.com/go-playground/validator"
	// "github.com/gosimple/slug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/gallery_go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	helper.PanicIfError(err)

	return db
}

func SetupRouter(db *gorm.DB) http.Handler {
	userRepository := repository.NewUserRepository()
	validate := validator.New()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)
	return router
}

func AddJWTToCookie(request *http.Request) {
	user := domain.User{
		ID:       1,
		Username: "test",
		FullName: "Test",
		Password: "password",
	}

	jwtToken, err := helper.CreateToken(user)
	if err != nil {
		helper.PanicIfError(err)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	request.AddCookie(cookie)
}

func TruncateTables(db *gorm.DB, tables ...string) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	for _, table := range tables {
		db.Exec("TRUNCATE TABLE " + table)
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}

func CreateUser(db *gorm.DB, username string) domain.User {
	hashedPassword, err := helper.HashPassword("test")
	helper.PanicIfError(err)

	user := domain.User{
		Username: username,
		FullName: "Test",
		Password: hashedPassword,
	}

	db.Transaction(func(tx *gorm.DB) error {
		var err error
		userRepository := repository.NewUserRepository()
		user, err = userRepository.Create(context.Background(), tx, user)
		helper.PanicIfError(err)

		return nil
	})

	return user
}

