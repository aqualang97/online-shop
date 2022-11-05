package repositories

import (
	"database/sql"
	"errors"
	"log"
	"online-shop/internal/models"
)

type UserRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetUserByID(ID int) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash, created_at, updated_at FROM users WHERE id=$1",
		ID).Scan(
		&user.ID, &user.Login, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash, created_at, updated_at FROM users WHERE id=($1)"+
		"", login).Scan(&user.ID, &user.Login, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash, created_at, updated_at FROM users WHERE id=($1)"+
		"", email).Scan(&user.ID, &user.Login, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) (int, error) {
	if user == nil {
		return 0, errors.New("user reg data is empty")
	}
	if r.TX != nil {
		prepare, err := r.TX.Prepare("INSERT INTO users( login, email, password_hash) VALUES ($1,$2,$3) RETURNING id")
		if err != nil {
			return 0, err
		}
		err = prepare.QueryRow(user.Login, user.Email, user.PasswordHash).Scan(&user.ID)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	}
	prepare, err := r.DB.Prepare("INSERT INTO users( login, email, password_hash) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		return 0, err

	}
	err = prepare.QueryRow(user.Login, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepo) UpdateUser(user *models.User) (int, error) {
	if user == nil {
		return 0, errors.New("user data is empty")
	}

	if r.TX != nil {
		prepare, err := r.TX.Prepare("UPDATE users SET login=$2, email=$3, password_hash=$4 WHERE id=$1")
		if err != nil {
			return 0, err
		}
		_, err = prepare.Exec(user.ID, user.Login, user.Email, user.PasswordHash)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	}
	prepare, err := r.TX.Prepare("UPDATE users SET login=$2, email=$3, password_hash=$4 WHERE id=$1")
	if err != nil {
		return 0, err
	}
	_, err = prepare.Exec(user.ID, user.Login, user.Email, user.PasswordHash)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
func (r *UserRepo) DeleteUser(id int) (int, error) {
	if r.TX != nil {
		_, err := r.TX.Exec("DELETE FROM users WHERE id=$1", id)
		if err != nil {
			return 0, err
		}
		return id, err
	}
	_, err := r.TX.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return id, err
}
func (r *UserRepo) BeginTX() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}
func (r *UserRepo) RollbackTX() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}
func (r *UserRepo) CommitTX() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}
func (r *UserRepo) SetTX(tx *sql.Tx) {
	r.TX = tx
}
func (r *UserRepo) GetTX() *sql.Tx {
	return r.TX
}
