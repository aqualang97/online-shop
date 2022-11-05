package controller

import "online-shop/internal/services"

type AuthController struct {
	services *services.Manager
}

func NewAuthController(services *services.Manager) *AuthController {
	return &AuthController{
		services: services,
	}
}
