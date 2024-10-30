package models

type Product struct {
	BaseModel
	Code        string  `gorm:"unique;not null" json:"code" form:"code"`
	Name        string  `gorm:"not null" json:"name" form:"name"`
	Description *string `json:"description" form:"description"`
	Price       float64 `gorm:"not null" json:"price" form:"price"`
	ImagePath   *string `json:"image_path"`
}
