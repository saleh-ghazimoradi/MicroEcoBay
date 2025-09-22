package domain

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID      uint    `gorm:"index" json:"cart_id"`
	ProductID   uint    `gorm:"index" json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Qty         uint    `json:"qty"`
	Price       float64 `json:"price"`
}

type Cart struct {
	gorm.Model
	UserId    uint       `gorm:"index" json:"user_id"`
	ProductID uint       `gorm:"uniqueIndex" json:"product_id"`
	Items     []CartItem `gorm:"constraint:OnDelete:CASCADE" json:"items"`
}

func (CartItem) TableName() string { return "cart_Items" }
