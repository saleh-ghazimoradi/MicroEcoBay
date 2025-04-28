package dto

type AuthResponse struct {
	UserId uint    `json:"user_id"`
	Email  string  `json:"email"`
	Exp    float64 `json:"exp"`
}
