package models

type Permission struct {
	BaseModel
	Name string `gorm:"unique;not null" json:"name"`
}
