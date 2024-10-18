package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Product struct {
    BaseModel
	Code string `gorm:"unique;not null" json:"code"`
    Name  string  `gorm:"not null" json:"name"`
	Description *string `json:"description"`
    Price float64 `gorm:"not null" json:"price"`
}
