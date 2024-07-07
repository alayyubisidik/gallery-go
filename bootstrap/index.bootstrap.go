package bootstrap

import (
	appconfig "gallery_go/configs/app_config"
	"gallery_go/database"
	"gallery_go/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	database.ConnectDatabase()

	app := gin.Default()

	routes.InitRoute(app)

	app.Run(appconfig.PORT)

}