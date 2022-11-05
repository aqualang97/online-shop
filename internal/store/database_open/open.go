package database_open

import (
	"database/sql"
	"fmt"
	"online-shop/config"
	"time"
)

func Open() (*sql.DB, error) {
	cfg := config.NewConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.PostgresUserName, cfg.PostgresPassword, cfg.PostgresDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.PostgresMaxIdleCons)
	db.SetMaxOpenConns(cfg.PostgresMaxOpenCons)
	db.SetConnMaxLifetime(time.Duration(cfg.PostgresConsMaxLifeTime) * time.Second)
	fmt.Println("Successfully connected to DB!")

	return db, nil
}
