package controller

import (
	"online-shop/config"
	"online-shop/internal/services"
)

type MenuController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewMenuController(services *services.Manager, cfg *config.Config) *MenuController {
	return &MenuController{
		services: services,
		cfg:      cfg,
	}
}
