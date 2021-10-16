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
	entities "mynamebvh.com/blog/internal/entities"
)

type SqlServer interface {
	DB() *gorm.DB
}

type sqlServer struct {
	db *gorm.DB
}

// ConnectDB connect to db
func ConnectDB() SqlServer {
	var err error
	var db *gorm.DB

	// Config logger database
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	//Get value env
	sqlName := config.GetEnv("SQL_NAME")
	dbHost := config.GetEnv("DB_HOST")
	dbPort := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	// Connect db
	dsn := fmt.Sprintf("%s://%s:%s?database=%s", sqlName, dbHost, dbPort, dbName)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Println(fmt.Sprintf("Error to loading Database %s", err))
		return nil
	}

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&entities.User{}, &entities.Category{}, &entities.Post{}, &entities.Tag{})
	// db.AutoMigrate(&entities.User{}, &entities.Category{}, &entities.Post{})
	fmt.Println("Database Migrated")

	return &sqlServer{
		db: db,
	}
}

func (c sqlServer) DB() *gorm.DB {
	return c.db
}
