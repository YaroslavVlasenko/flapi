package models

import (
	"database/sql"
	"time"
)

// Locale struct
type Locale struct {
	ID        uint         `json:"id" faker:"-"`
	Name      string       `json:"name"`
	Default   int          `json:"default"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-" faker:"-"`
}
