package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	*gorm.Model
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type PermissionScope interface {
	GetScope() string
}
