package test

import (
	controller "gallery_go/controllers"
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouteTest(app *gin.Engine) *gin.Engine {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", controller.SignUp)
	route.POST("/api/v1/users/signin", controller.SignIn)

	return route
}	

func DeleteTestUsernames(db *gorm.DB) {
    // Hapus semua username yang berawalan "test"
    db.Exec("DELETE FROM users WHERE username LIKE 'test%'")
}

func CreateUser(username string, email string) models.User {
	hashedPassword, err := helper.HashPassword("test")
	helper.PanicIfError(err)

	user := models.User{
		Username: username,
		FullName: "Test",
		Email: email,
		Password: hashedPassword,
	}

	err = database.DB.Save(&user).Error
	helper.PanicIfError(err)
	
	return user
}
