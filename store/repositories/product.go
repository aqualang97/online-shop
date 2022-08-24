package repositories

import "database/sql"

type ProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}
