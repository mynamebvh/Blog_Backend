package db

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
	var db *gorm.DB
	_ = db
	
	port := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")
	dsn := "sqlserver://localhost:"+ port +"?database=" + dbName
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&model.User{},&model.Tag{}, &model.Category{})
	fmt.Println("Database Migrated")
}
