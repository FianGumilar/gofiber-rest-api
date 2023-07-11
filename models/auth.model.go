package models

type Login struct {
	Email    string `json:"email" form:"email" validate:"email,required" gorm:"not null"`
	Password string `json:"password" form:"password" validate:"required,gte=8"`
}
