package router

import (
	"github.com/gin-gonic/gin"
	"online-shop/config"
	"online-shop/internal/controller"
	"online-shop/internal/services"
)

func Router(services *services.Manager, cfg *config.Config, r *gin.Engine) {
	//mw := middleware.NewMiddleware(service)
	ctr := controller.NewController(services, cfg)

	r.POST("/login", func(c *gin.Context) { ctr.Auth.Login(c) })
	r.POST("/logout", func(c *gin.Context) { ctr.Auth.Logout(c) })
	r.POST("/registration", func(c *gin.Context) { ctr.Auth.Registration(c) })
	r.POST("/refresh", func(c *gin.Context) { ctr.Auth.Refresh(c) })

}
