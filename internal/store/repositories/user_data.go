package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"online-shop/internal/models"
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
	if r.TX != nil {
		prepare, err := r.TX.Prepare("INSERT INTO users_data(user_id, full_name, date_of_birth, number, address, discount_percent, discount_amount) VALUES ($1,$2,$3,$4,$5,$6,$7)")
		if err != nil {
			return uuid.Nil, err
		}
		_, err = prepare.Exec(userUIID, data.FullName, data.DateOfBirth, data.Number, data.Address, data.DiscountPercent, data.DiscountAmount)
		if err != nil {
			return uuid.Nil, err
		}
		return data.UserID, nil
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
