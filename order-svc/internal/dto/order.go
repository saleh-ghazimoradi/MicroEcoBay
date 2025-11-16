package dto

type CreateOrderRequest struct {
	UserId uint `json:"user_id"`
	Items  []struct {
		ProductId uint    `json:"product_id"`
		Quantity  uint    `json:"quantity"`
		Price     float64 `json:"price"`
	} `json:"items"`
}

type CreateOrderResponse struct {
	OrderId  uint    `json:"order_id"`
	OrderRef string  `json:"order_ref"`
	Total    float64 `json:"total"`
	Status   string  `json:"status"`
}
