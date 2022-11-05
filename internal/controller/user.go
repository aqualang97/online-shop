package controller

import (
	"online-shop/config"
	"online-shop/internal/services"
)

type UserController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewUserController(services *services.Manager, cfg *config.Config) *UserController {
	return &UserController{
		services: services,
		cfg:      cfg,
	}
}
