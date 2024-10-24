package models

type Product struct {
	BaseModel
	Code        string  `gorm:"unique;not null" json:"code"`
	Name        string  `gorm:"not null" json:"name"`
	Description *string `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}
