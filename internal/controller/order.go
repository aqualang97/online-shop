package controller

import (
	"online-shop/config"
	"online-shop/internal/services"
)

type OrderController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewOrderController(services *services.Manager, cfg *config.Config) *OrderController {
	return &OrderController{
		services: services,
		cfg:      cfg,
	}
}
