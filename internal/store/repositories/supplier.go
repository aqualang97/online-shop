package repositories

import "database/sql"

type SupplierRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewSupplierRepo(db *sql.DB) *SupplierRepo {
	return &SupplierRepo{DB: db}
}
