package handlers

import "github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
