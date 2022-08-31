package repositories

import (
	"app/internal/models"
	"database/sql"
	"github.com/google/uuid"
)

type UserDataRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserDataRepo(db *sql.DB) *UserDataRepo {
	return &UserDataRepo{DB: db}
}

func (r *UserDataRepo) CreateUserDate(data *models.UserData) (uuid.UUID, error) {
	if data == nil {
		
	}
	userUIID, err := data.UserID.MarshalBinary()
	if err != nil {
		return uuid.Nil, err
	}

	return data.UserID, nil
}
func (r *UserDataRepo) UpdateUserDate() {
	//if we will upd info, we should save in memory data about discount
}
func (r *UserDataRepo) DeleteUserDate() {

}
func (r *UserDataRepo) AddDiscountAmount() {

}
func (r *UserDataRepo) UpdateDiscountPercent() {

}
func (r *UserDataRepo) CheckUserData() {

}

func (r *UserDataRepo) BeginTX() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}
func (r *UserDataRepo) RollbackTX() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}
func (r *UserDataRepo) CommitTX() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}
func (r *UserDataRepo) SetTX(tx *sql.Tx) {
	r.TX = tx
}
func (r *UserDataRepo) GetTX() *sql.Tx {
	return r.TX
}
