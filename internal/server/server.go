package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online-shop/config"
	"online-shop/internal/router"
	"online-shop/internal/services"
	"online-shop/internal/store"
	"sync"
)

type Server struct {
	cfg *config.Config
	db  *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Start() {
	storage := store.NewStore(s.db)

	service, err := services.NewManager(storage)
	if err != nil {
		log.Fatal(err)
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		// Run Online-shop
		routerUser := gin.Default()

		router.UserRouter(service, s.cfg, routerUser)

		err := http.ListenAndServe("localhost:8000", routerUser)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()

	}()
	go func() {
		//Run images server
		routerAdmin := gin.Default()
		router.AdminRouter(service, s.cfg, routerAdmin)

		//router.Router(service, s.cfg, routerAdmin)

		err := http.ListenAndServe("localhost:8080", routerAdmin)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
