package database

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/Rene-Michel99/API-GO-REST/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var DATABASE DBinstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if error != nil {
		log.Fatal("Error trying to connect to postgres.\n", error)
		os.Exit(2)
	}
	log.Println("Successfully connected to postgres!")
	db.Logger = newLogger

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Book{})

	DATABASE = DBinstance{DB: db}
}


func Insert(book *models.Book) *gorm.DB {
    return DATABASE.DB.Create(&book)
}

func Get(book *models.Book) *gorm.DB {
    return DATABASE.DB.Where("id = ? OR name = ?", book.ID, book.Name).Find(&book)
}

func Update(book *models.Book) *gorm.DB {
	return DATABASE.DB.Model(&models.Book{}).Where("id = ?", book.ID).Update("CopyQnt", book.CopyQnt)
}

func Delete(book *models.Book) *gorm.DB {
	return DATABASE.DB.Delete(book)
}
