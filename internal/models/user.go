package models

import (
	"database/sql"
	"time"
)

//User struct
type User struct {
	ID          uint         `json:"id"`
	Username    string       `json:"username"`
	Email       string       `json:"email"`
	Roles       string       `json:"-"`
	Password    string       `json:"-"`
	LocaleID    *uint        `json:"-"`
	Locale      Locale       `json:"locale" gorm:"foreignKey:LocaleID"`
	ActivatedAt sql.NullTime `json:"-"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"-"`
}
