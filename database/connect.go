package database

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	config "mynamebvh.com/blog/config"
	model "mynamebvh.com/blog/models"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	port := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	dsn := "sqlserver://localhost:"+ port +"?database=" + dbName
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{},&model.Category{}, &model.Post{},&model.Tag{})
	fmt.Println("Database Migrated")
}
