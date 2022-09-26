package repositories

import "database/sql"

type TokenRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewTokenRepo(db *sql.DB) *TokenRepo {
	return &TokenRepo{DB: db}
}
