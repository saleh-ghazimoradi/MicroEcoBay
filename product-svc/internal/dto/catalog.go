package dto

type CreateCategoryReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"imageUrl" validate:"required"`
}

type UpdateCategoryReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type CreateProductReq struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CategoryID  uint    `json:"categoryId" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       uint    `json:"stock" validate:"required"`
	ImageURL    string  `json:"imageUrl" validate:"required"`
	Status      string  `json:"status" validate:"required"`
}

type UpdateProductReq struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *uint    `json:"stock"`
	ImageURL    *string  `json:"imageUrl"`
	Status      *string  `json:"status"`
}
