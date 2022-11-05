package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	PostgresUserName           string
	PostgresPassword           string
	PostgresDBName             string
	Host                       string
	Port                       int
	PostgresMaxIdleCons        int
	PostgresMaxOpenCons        int
	PostgresConsMaxLifeTime    int
	AccessTokenLifeTime        int
	RefreshTokenLifeTime       int
	AccessSecret               string
	RefreshSecret              string
	AccessControlAllowOrigin   string
	AccessControlAllowHeaders  string
	AccessControlExposeHeaders string
	AccessControlAllowMethods  string
}

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load("config/postgres.env")
	if err != nil {
		log.Fatal("Can`t load postgres.env")
	}
	cfg.PostgresUserName = os.Getenv("user")
	cfg.PostgresPassword = os.Getenv("pwd")
	cfg.PostgresDBName = os.Getenv("dbName")
	cfg.Host = os.Getenv("host")
	cfg.Port, _ = strconv.Atoi(os.Getenv("port"))
	cfg.PostgresMaxIdleCons, _ = strconv.Atoi(os.Getenv("maxIdleConns"))
	cfg.PostgresMaxOpenCons, _ = strconv.Atoi(os.Getenv("maxOpenConns"))
	cfg.PostgresConsMaxLifeTime, _ = strconv.Atoi(os.Getenv("connMaxLifetime"))

	err = godotenv.Load("config/token.env")
	if err != nil {
		log.Fatal("Cant load token.env")
	}
	cfg.AccessTokenLifeTime, _ = strconv.Atoi(os.Getenv("AccessTokenLifeTime"))
	cfg.RefreshTokenLifeTime, _ = strconv.Atoi(os.Getenv("RefreshTokenLifeTime"))
	cfg.AccessSecret = os.Getenv("AccessSecret")
	cfg.RefreshSecret = os.Getenv("RefreshSecret")

	err = godotenv.Load("config/cors.env")
	if err != nil {
		log.Fatal("Cant load cors.env")
	}

	cfg.AccessControlAllowOrigin = os.Getenv("AccessControlAllowOrigin")
	cfg.AccessControlAllowHeaders = os.Getenv("AccessControlAllowHeaders")
	cfg.AccessControlExposeHeaders = os.Getenv("AccessControlExposeHeaders")
	cfg.AccessControlAllowMethods = os.Getenv("AccessControlAllowMethods")
	return cfg
}
