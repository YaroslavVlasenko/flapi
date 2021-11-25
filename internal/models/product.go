package models

import (
	"database/sql"
	"time"
)

// Product struct
type Product struct {
	ID           uint                 `json:"id"`
	Amount       int                  `json:"amount"`
	Price        float64              `json:"price"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	DeletedAt    sql.NullTime         `json:"-"`
	Translations []ProductTranslation `json:"translations" gorm:"foreignKey:TranslatableID;references:ID"`
}
