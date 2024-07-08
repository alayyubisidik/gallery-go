package test

import (
	controller "gallery_go/controllers"
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouteTest(app *gin.Engine) *gin.Engine {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", controller.SignUp)
	route.POST("/api/v1/users/signin", controller.SignIn)
	route.GET("/api/v1/users/currentuser", controller.CurrentUser)

	return route
}	

func DeleteTestUsernames(db *gorm.DB) {
    // Hapus semua username yang berawalan "test"
    db.Exec("DELETE FROM users WHERE username LIKE 'test%'")
}

func AddJWTToCookie(request *http.Request) {
	user := models.User{
		ID:       1,
		Username: "test",
		FullName: "Test",
		Email:    "test@gmail.com",
		Role: "author",
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
