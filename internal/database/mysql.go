package database

import (
	"errors"
	"fmt"
	"github.com/YaroslavVlasenko/flapi/internal/configs"
	"github.com/YaroslavVlasenko/flapi/internal/models"
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

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
			locales := []models.Locale{
				{Name: "uk_UA", Default: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{Name: "ru_RU", Default: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{Name: "en_GB", Default: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			}
			DB.Create(&locales)

			for i := 0; i < 30; i++ {
				a := models.Product{}
				err := faker.FakeData(&a)
				if err != nil {
					return
				}
				DB.Create(&a)
			}

			for i := 0; i < 30; i++ {
				a := models.Country{}
				err := faker.FakeData(&a)
				if err != nil {
					return
				}
				DB.Create(&a)
			}

			for i := 0; i < 30; i++ {
				a := models.Brand{}
				err := faker.FakeData(&a)
				if err != nil {
					return
				}
				DB.Create(&a)
			}

			//products := []models.Product{
			//	{Amount: 2, Price: 200.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{Amount: 4, Price: 300.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{Amount: 6, Price: 100.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//}
			//DB.Create(&products)
			//
			//productTranslations := []models.ProductTranslation{
			//	{TranslatableID: 1, Title: "??????????", Description: "????????", LocaleID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 1, Title: "????????????????", Description: "????????????????", LocaleID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 1, Title: "Title", Description: "Description", LocaleID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 2, Title: "?????????? 2", Description: "???????? 2", LocaleID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 2, Title: "???????????????? 2", Description: "???????????????? 2", LocaleID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 2, Title: "Title 2", Description: "Description 2", LocaleID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 3, Title: "???????????? 42", Description: "?????????????? 42", LocaleID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 3, Title: "???????????? 42", Description: "?????????????? 42", LocaleID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//	{TranslatableID: 3, Title: "Qwerty 42", Description: "Asdfghj 42", LocaleID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			//}
			//DB.Create(&productTranslations)
		}
	}
}
