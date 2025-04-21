package domain

import "gorm.io/gorm"

type Address struct {
	ID           uint   `gorm:"PrimaryKey" json:"id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostCode     string `json:"post_code"`
	Country      string `json:"country"`
	UserId       uint   `json:"user_id"`
	gorm.Model
}
