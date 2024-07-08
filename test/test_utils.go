package test

import (
	"gorm.io/gorm"
	controller "gallery_go/controllers"
	"gallery_go/exception"
	"github.com/gin-gonic/gin"

)

func InitRouteTest(app *gin.Engine) *gin.Engine {
	route := app

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", controller.SignUp)

	return route
}	

func DeleteTestUsernames(db *gorm.DB) {
    // Hapus semua username yang berawalan "test"
    db.Exec("DELETE FROM users WHERE username LIKE 'test%'")
}