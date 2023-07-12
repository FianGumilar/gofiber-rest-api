package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" form:"title" validate:"required" gorm:"not null"`
	Author    string         `json:"author" form:"author" validate:"required" gorm:"not null"`
	Cover     string         `json:"cover" form:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
