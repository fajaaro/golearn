package models

type UserPermission struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	PermissionID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_permission_user,unique;" json:"permission_id"`
	UserID       uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_permission_user,unique;" json:"user_id"`

	Permission Permission `json:"-"`
	User       User       `json:"-"`
}