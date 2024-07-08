package bootstrap

import (
	"gallery_go/database"
	"gallery_go/helper"
	"gallery_go/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	database.ConnectDatabase()

	app := gin.Default()

	routes.InitRoute(app)

	app.Run(":" + os.Getenv("PORT"))

}