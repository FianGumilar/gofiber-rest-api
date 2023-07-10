package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" form:"name" validate:"min=6,max=32" gorm:"not null"`
	Email     string         `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Phone     int            `json:"phone" form:"phone" validate:"required,number,min=10" gorm:"required,not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
