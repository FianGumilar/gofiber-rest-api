package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Photo      string         `json:"photos" form:"photos" validate:"required" gorm:"not null"`
	CategoryId uint           `json:"category_id" form:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
