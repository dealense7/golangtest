package models

type Role struct {
	Name        string       `json:"name" gorm:"size:255"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	BaseModel
}
