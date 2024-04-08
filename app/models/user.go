package models

const AdminEmail = "j.doe@example.com"

type User struct {
	FistName    string       `json:"firstName" gorm:"size:255"`
	LastName    string       `json:"lastName" gorm:"size:255"`
	Email       string       `json:"email" gorm:"unique;not null;size:255"`
	Password    string       `json:"-" gorm:"not null"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	Roles       []Role       `gorm:"many2many:user_roles;"`
	BaseModel
}

func (u User) GetScope() string {
	return "users"
}

func (u *User) HasPermission(permission string) bool {
	return true
}
