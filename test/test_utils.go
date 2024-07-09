package test

import (
	usercontroller "gallery_go/controllers/user_controller"
	"gallery_go/database"
	"gallery_go/exception"
	"gallery_go/helper"
	"gallery_go/middleware"
	"gallery_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouteTest(app *gin.Engine) *gin.Engine {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", usercontroller.SignUp)
	route.POST("/api/v1/users/signin", usercontroller.SignIn)
	route.GET("/api/v1/users/currentuser", usercontroller.CurrentUser)
	route.DELETE("/api/v1/users/signout", middleware.AuthMidddleware, usercontroller.SignOut)
	route.PUT("/api/v1/users/:userId", middleware.AuthMidddleware, usercontroller.Update)

	return route
}	

func DeleteTestUsernames(db *gorm.DB) {
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
