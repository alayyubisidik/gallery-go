package routes

import (
	usercontroller "gallery_go/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.POST("/api/v1/users/signup", usercontroller.SignUp)
}	