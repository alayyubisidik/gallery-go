package app

import (
	"gallery_go/helper"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	dsn := config.GetString("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	return db
}

// migrate create -ext sql -dir db/migrations create_users_table

// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations down
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations force 9863498326134
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/gallery_go?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations version