package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"online-shop/config"
	"online-shop/internal/server"
	"online-shop/internal/store/database_open"
)

func main() {
	db, err := database_open.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cfg := config.NewConfig()
	s := server.NewServer(cfg, db)

	s.Start()
	fmt.Println("Successfully connected!")

	//ud := repositories.NewUserDataRepo(database_open)
	//u := repositories.NewUserRepo(database_open)

}
