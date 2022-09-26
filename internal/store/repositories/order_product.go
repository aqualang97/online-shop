package repositories

import "database/sql"

type OrderProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderProductRepo(db *sql.DB) *OrderProductRepo {
	return &OrderProductRepo{DB: db}
}
