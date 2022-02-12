package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	U_email    string `gorm:"size:50;unique"  json:"u_email" form:"u_email" validate:"required,email"`
	U_password string `json:"u_password" form:"u_password"`
}
