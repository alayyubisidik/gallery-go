package routes

import (
	imagecontroller "gallery_go/controllers/image_controller"
	usercontroller "gallery_go/controllers/user_controller"
	"gallery_go/exception"
	"gallery_go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.Static("/public", "./public")

	route.Use(exception.GlobalErrorHandler())
	route.POST("/api/v1/users/signup", usercontroller.SignUp)
	route.POST("/api/v1/users/signin", usercontroller.SignIn)
	route.GET("/api/v1/users/currentuser", usercontroller.CurrentUser)
	route.DELETE("/api/v1/users/signout", middleware.AuthMidddleware, usercontroller.SignOut)
	route.PUT("/api/v1/users/:userId", middleware.AuthMidddleware, usercontroller.Update)

	route.POST("/api/v1/images", middleware.AuthMidddleware, imagecontroller.Store)
	route.DELETE("/api/v1/images/:imageId", middleware.AuthMidddleware, imagecontroller.Delete)

}	
