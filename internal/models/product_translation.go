package models

import (
	"database/sql"
	"time"
)

// ProductTranslation struct for translation fields
type ProductTranslation struct {
	ID             uint         `json:"id"`
	TranslatableID uint         `json:"-"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	LocaleID       uint         `json:"-"`
	Locale         Locale       `json:"locale,omitempty" gorm:"foreignKey:LocaleID"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"-"`
}
