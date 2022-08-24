package repositories

import "database/sql"

type UserDataRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserDataRepo(db *sql.DB) *UserDataRepo {
	return &UserDataRepo{DB: db}
}
