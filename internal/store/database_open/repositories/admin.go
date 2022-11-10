package repositories

import (
	"database/sql"
)

type AdminRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func (a *AdminRepo) CheckRole(login string) (bool, error) {
	var role string
	if a.TX != nil {
		err := a.TX.QueryRow("SELECT role FROM users WHERE login=$1", login).Scan(&role)
		if err != nil {
			return false, err
		}
	}
	if role == "user" {
		return false, nil
	}
	return true, nil

}
func (a *AdminRepo) MakeRole(login, role string) (string, error) {

	if a.TX != nil {
		_, err := a.TX.Exec("UPDATE users SET role=$2 WHERE login=$1", login, role)
		if err != nil {
			return login, err
		}
		return login, err
	}
	_, err := a.TX.Exec("UPDATE users SET role=$2 WHERE login=$1", login, role)
	if err != nil {
		return login, err
	}
	return login, nil

}

func (a *AdminRepo) BeginTX() error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	a.TX = tx
	return nil
}
func (a *AdminRepo) RollbackTX() error {
	defer func() {
		a.TX = nil
	}()
	if a.TX != nil {
		return a.TX.Rollback()
	}
	return nil
}
func (a *AdminRepo) CommitTX() error {
	defer func() {
		a.TX = nil
	}()
	if a.TX != nil {
		return a.TX.Commit()
	}
	return nil
}
func (a *AdminRepo) SetTX(tx *sql.Tx) {
	a.TX = tx
}
func (a *AdminRepo) GetTX() *sql.Tx {
	return a.TX
}
