package models

type UserRole struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_user_role,unique;" json:"user_id"`
	RoleID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_user_role,unique;" json:"role_id"`

	User User `json:"-"`
	Role Role `json:"-"`
}