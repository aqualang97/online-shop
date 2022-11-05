package router

import (
	"github.com/gin-gonic/gin"
	"online-shop/config"
	"online-shop/internal/services"
)

func Router(services *services.Manager, cfg *config.Config, r *gin.Engine) {
	//mw := middleware.NewMiddleware(service)
	//ctr := controller.NewController(services, cfg)

}
