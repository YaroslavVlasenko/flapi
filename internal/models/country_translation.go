package models

import (
	"database/sql"
	"time"
)

// CountryTranslation struct struct for translation fields
type CountryTranslation struct {
	ID             uint         `json:"id"`
	TranslatableID uint         `json:"-"`
	Title          string       `json:"title"`
	LocaleID       uint         `json:"-"`
	Locale         Locale       `json:"locale" gorm:"foreignKey:LocaleID"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"-"`
}
