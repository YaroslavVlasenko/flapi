package database

import (
	"app-backend/internal/configs"
	"app-backend/internal/models"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// DB gorm connector
var DB *gorm.DB

// Connect connect to db
func Connect() {
	var err error

	p := configs.Load("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.Load("DB_USER"),
		configs.Load("DB_PASSWORD"),
		configs.Load("DB_HOST"),
		port,
		configs.Load("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}

// AutoMigrate apply migrations
func AutoMigrate() {
	if err := DB.AutoMigrate(
		models.Locale{},
		models.Product{},
		models.ProductTranslation{},
		models.User{},
		models.Brand{},
		models.Country{},
		models.CountryTranslation{}); err == nil && DB.Migrator().HasTable(&models.Locale{}) {
		if err := DB.First(&models.Locale{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&models.Locale{Name: "Uk-ua", Default: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
			DB.Create(&models.Locale{Name: "Ru-ru", Default: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()})

			DB.Create(&models.Product{Amount: 3, Price: 200.0, CreatedAt: time.Now(), UpdatedAt: time.Now()})
			DB.Create(&models.Product{Amount: 2, Price: 400.0, CreatedAt: time.Now(), UpdatedAt: time.Now()})

			DB.Create(&models.ProductTranslation{TranslatableID: 1, Title: "Title", Description: "Description", LocaleID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
			DB.Create(&models.ProductTranslation{TranslatableID: 1, Title: "Название", Description: "Описание", LocaleID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()})
			DB.Create(&models.ProductTranslation{TranslatableID: 2, Title: "Qwerty", Description: "Asdfghj", LocaleID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		}
	}
}
