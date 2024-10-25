package models

type RolePermission struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	PermissionID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_permission_role,unique;" json:"permission_id"`
	RoleID       uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;index:idx_permission_role,unique;" json:"role_id"`

	Permission Permission `json:"-"`
	Role       Role       `json:"-"`
}