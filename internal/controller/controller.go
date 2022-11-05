package controller

import (
	"online-shop/config"
	"online-shop/internal/services"
)

type Controller struct {
	Auth  *AuthController
	User  *UserController
	Menu  *MenuController
	Order *OrderController
}

func NewController(services *services.Manager, cfg *config.Config) *Controller {
	return &Controller{
		Auth:  NewAuthController(services),
		User:  NewUserController(services, cfg),
		Menu:  NewMenuController(services, cfg),
		Order: NewOrderController(services, cfg),
	}
}
