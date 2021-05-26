package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	U_full_name string `gorm:"size:100" json:"u_full_name" form:"u_full_name" `
	U_tel       string `gorm:"size:20" json:"u_tel" form:"u_tel"`
	U_address   string `gorm:"size:100" json:"u_address" form:"u_address"`
	U_email     string `gorm:"size:50;unique"  json:"u_email" form:"u_email" validate:"required,email"`
	U_password  string `json:"u_password" form:"u_password"`
}
