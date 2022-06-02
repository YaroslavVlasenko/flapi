package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Product struct
type Product struct {
	ID           uint                 `json:"id"`
	Amount       int                  `json:"amount"`
	Price        float32              `json:"price"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	DeletedAt    sql.NullTime         `json:"-"`
	Translations []ProductTranslation `json:"translations" gorm:"foreignKey:TranslatableID;references:ID"`
}

type UpdateProductInput struct {
	Amount       int                  `json:"amount"`
	Price        float32              `json:"price"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}
func (p *Product) AfterSave(tx *gorm.DB) (err error) {
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
func (p *Product) AfterUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}
func (p *Product) AfterDelete(tx *gorm.DB) (err error) {
	return nil
}