package dto

type Cart struct {
	ProductID   uint    `json:"product_id" validate:"required"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Qty         uint    `json:"qty" validate:"required,gt=0"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}

type CartUpdateQty struct {
	ProductID uint `json:"product_id" validate:"required"`
	Qty       uint `json:"qty" validate:"required,gte=0"`
}

type CartRemove struct {
	ProductID uint `json:"product_id" validate:"required"`
}

type CartItem struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Qty         uint    `json:"qty"`
	Price       float64 `json:"price"`
	LineTotal   float64 `json:"line_total"`
}

type CartResponse struct {
	UserID   uint       `json:"user_id"`
	Items    []CartItem `json:"items"`
	Count    int        `json:"count"`
	SubTotal float64    `json:"sub_total"`
}
