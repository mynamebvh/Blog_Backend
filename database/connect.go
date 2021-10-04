package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	config "mynamebvh.com/blog/config"
	model "mynamebvh.com/blog/models"
)

// ConnectDB connect to db
func ConnectDB() {

	newLogger := logger.New(
  log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
  logger.Config{
    SlowThreshold:              time.Second,   // Slow SQL threshold
    LogLevel:                   logger.Silent, // Log level
    IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
    Colorful:                  false,          // Disable color
  },
)

	var err error
	port := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	dsn := "sqlserver://localhost:"+ port +"?database=" + dbName
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{},&model.Category{}, &model.Post{},&model.Tag{})
	fmt.Println("Database Migrated")
}
