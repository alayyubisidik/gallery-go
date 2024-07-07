package database

import (
	"fmt"
	dbconfig "gallery_go/configs/db_config"
	"gallery_go/helper"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_HOST, dbconfig.DB_PORT, dbconfig.DB_NAME)

	DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(errConnection)

	log.Println("Connected to database")
}

// migrate create -ext sql -dir database/migrations create_users_table

// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations up
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations down
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations force 9863498326134
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path database/migrations version