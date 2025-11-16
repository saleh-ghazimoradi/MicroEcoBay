package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId     uint        `json:"user_id"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"`
	OrderRef   string      `json:"order_ref"`
	OrderItems []OrderItem `json:"order_items" gorm:"constraint:OnDelete:CASCADE;foreignKey:OrderId"`
}

type OrderItem struct {
	gorm.Model
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}
