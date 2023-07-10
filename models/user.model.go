package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" form:"name" gorm:"not null"`
	Email     string         `json:"email" form:"email" validate:"email" gorm:"not null"`
	Phone     string         `json:"phone" form:"phone" gorm:"required,not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
