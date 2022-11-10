package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online-shop/config"
	"online-shop/internal/models"
	"online-shop/internal/services"
)

type Middleware struct {
	service *services.Manager
}

func NewMiddleware(service *services.Manager) *Middleware {
	return &Middleware{
		service: service,
	}
}

func (m *Middleware) SetCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := config.NewConfig()
		ctx.Request.Header.Set("Access-Control-Allow-Origin", cfg.AccessControlAllowOrigin)
		ctx.Request.Header.Set("Access-Control-Allow-Methods", cfg.AccessControlAllowMethods)
		ctx.Request.Header.Set("Access-Control-Allow-Headers", cfg.AccessControlAllowHeaders)
		ctx.Request.Header.Set("Access-Control-Expose-Headers", cfg.AccessControlExposeHeaders)
		ctx.Next()
	}
}
func (m *Middleware) IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := models.UserLoginRequest{}
		if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		admin, err := m.service.Admin.IsAdmin(req.Login)
		if err != nil {
			log.Println(err)
			return
		}
		if !admin {
			http.Error(ctx.Writer, errors.New("you are not admin").Error(), http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
