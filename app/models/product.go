package models

type Product struct {
	BaseModel
	Code        string  `gorm:"unique;not null" json:"code" form:"code" excel-col-index:"0"`
	Name        string  `gorm:"not null" json:"name" form:"name" excel-col-index:"1"`
	Description *string `json:"description" form:"description" excel-col-index:"2"`
	Price       float64 `gorm:"not null" json:"price" form:"price" excel-col-index:"3"`
	ImagePath   *string `json:"image_path"`
}
