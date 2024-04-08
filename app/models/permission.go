package models

type Permission struct {
	Name        string `json:"name" gorm:"size:255"`
	Description string `json:"description" gorm:"size:255"`
	BaseModel
}
