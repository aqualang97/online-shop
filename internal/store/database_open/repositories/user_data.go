package repositories

import (
	"database/sql"
	"errors"
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
		return uuid.Nil, errors.New("user data is empty")
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
	prepare, err := r.DB.Prepare("INSERT INTO users_data(user_id, full_name, date_of_birth, number, address, discount_percent, discount_amount) VALUES ($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		return uuid.Nil, err
	}
	_, err = prepare.Exec(userUIID, data.FullName, data.DateOfBirth, data.Number, data.Address, data.DiscountPercent, data.DiscountAmount)
	if err != nil {
		return uuid.Nil, err
	}
	return data.UserID, nil
}
func (r *UserDataRepo) UpdateUserDate(data *models.UserData) (*models.UserData, error) {
	//if we will upd info, we should save in memory data about discount

	if data == nil {
		return nil, errors.New("user data is empty")
	}
	//userUIID, err := data.UserID.MarshalBinary()
	//if err != nil {
	//	return nil, err
	//}
	if r.TX != nil {
		prepare, err := r.TX.Prepare("UPDATE users_data SET (full_name=$2, date_of_birth=$3, number=$4, address=$5) WHERE user_id=$1")
		if err != nil {
			return nil, err
		}
		_, err = prepare.Exec(data.UserID, data.FullName, data.DateOfBirth, data.Number, data.Address)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
	prepare, err := r.TX.Prepare("UPDATE users_data SET (full_name=$2, date_of_birth=$3, number=$4, address=$5) WHERE user_id=$1")
	if err != nil {
		return nil, err
	}
	_, err = prepare.Exec(data.UserID, data.FullName, data.DateOfBirth, data.Number, data.Address)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *UserDataRepo) DeleteUserDate(userID int) error {
	if r.TX != nil {
		_, err := r.TX.Exec("DELETE FROM users_data WHERE user_id=$1", userID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := r.TX.Exec("DELETE FROM users_data WHERE user_id=$1", userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserDataRepo) AddDiscountAmount(discount float32, data *models.UserData) error {
	// Cashback
	if r.TX != nil {
		prepare, err := r.TX.Prepare("UPDATE users_data SET discount_amount=discount_amount+$2 WHERE user_id=$1")
		if err != nil {
			return err
		}
		_, err = prepare.Exec(data.UserID, discount)
		if err != nil {
			return err
		}

		return nil
	}
	prepare, err := r.DB.Prepare("UPDATE users_data SET discount_amount=discount_amount+$2 WHERE user_id=$1")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(data.UserID, discount)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserDataRepo) UpdateDiscountPercent(discountPercent int, data *models.UserData) error {

	if r.TX != nil {
		prepare, err := r.TX.Prepare("UPDATE users_data SET discount_percent=discount_percent+$2 WHERE user_id=$1")
		if err != nil {
			return err
		}
		_, err = prepare.Exec(data.UserID, discountPercent)
		if err != nil {
			return err
		}

		return nil
	}
	prepare, err := r.DB.Prepare("UPDATE users_data SET discount_percent=discount_percent+$2 WHERE user_id=$1")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(data.UserID, discountPercent)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserDataRepo) GetUserData(id int) (*models.UserData, error) {
	var data models.UserData
	err := r.DB.QueryRow("SELECT user_id, full_name, date_of_birth, number, address, discount_percent, "+
		"discount_amount, created_at, updated_at FROM users_data WHERE user_id=($1)", id).Scan(
		&data.UserID, &data.FullName, &data.DateOfBirth, &data.Number, &data.Address,
		&data.DiscountPercent, &data.DiscountAmount, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return &data, err
	}

	return &data, nil
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
