package domain

import "gorm.io/gorm"

type Category struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	Status      string `json:"status"`
	gorm.Model
}
