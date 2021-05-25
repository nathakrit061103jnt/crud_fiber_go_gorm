package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	P_Name            string    `gorm:"size:100" json:"p_name" form:"p_name"`
	P_Price           float64   `json:"p_price" form:"p_price"`
	P_Date            time.Time `json:"p_date" form:"p_date"`
	P_Amount          int       `json:"p_amount" form:"p_amount"`
	P_Expiration_Date time.Time `json:"p_expiration_date" form:"p_expiration_date"`
}
