package models

import (
	"database/sql"
	"time"
)

// Brand struct
type Brand struct {
	ID        uint         `json:"id"`
	Code      string       `json:"code"`
	Thumbnail string       `json:"thumbnail"`
	IsPublish int          `json:"is_publish"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-"`
}
