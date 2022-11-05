package middleware

import (
	"net/http"
	"online-shop/config"
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

func (m *Middleware) SetCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig()
		w.Header().Set("Access-Control-Allow-Origin", cfg.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", cfg.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", cfg.AccessControlAllowHeaders)
		w.Header().Set("Access-Control-Expose-Headers", cfg.AccessControlExposeHeaders)
		next.ServeHTTP(w, r)
	})
}
