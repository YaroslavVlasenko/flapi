package models

import (
	"database/sql"
	"time"
)

// Country struct
type Country struct {
	ID           uint                 `json:"id"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	DeletedAt    sql.NullTime         `json:"-"`
	Translations []CountryTranslation `json:"translations" gorm:"foreignKey:TranslatableID;references:ID;"`
}
