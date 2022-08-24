package repositories

import (
	"app/internal/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
)

type UserRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetUserByID(ID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash created_at, updated_at FROM users WHERE id=($1)", ID).Scan(user.ID, user.Login, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash created_at, updated_at FROM users WHERE id=($1)", login).Scan(user.ID, user.Login, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, login, email, password_hash created_at, updated_at FROM users WHERE id=($1)", email).Scan(user.ID, user.Login, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) (uuid.UUID, error) {
	if user == nil {
		return uuid.Nil, errors.New("user data is empty")
	}
	userUIID, err := user.ID.MarshalBinary()
	if err != nil {
		return uuid.Nil, err
	}
	if r.TX != nil {
		prepare, err := r.TX.Prepare("INSERT INTO users(id, login, email, password_hash) VALUES ($1,$2,$3,$4)")
		if err != nil {
			return uuid.Nil, err
		}
		_, err = prepare.Exec(userUIID, user.Login, user.Email, user.PasswordHash)
		if err != nil {
			return uuid.Nil, err
		}
		return user.ID, nil
	}
	prepare, err := r.DB.Prepare("INSERT INTO users(id, login, email, password_hash) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return uuid.Nil, err
	}
	_, err = prepare.Exec(userUIID, user.Login, user.Email, user.PasswordHash)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

func (r *UserRepo) UpdateUser(user *models.User) (uuid.UUID, error) {
	if user == nil {
		return uuid.Nil, errors.New("user data is empty")
	}
	userUIID, err := user.ID.MarshalBinary()
	if err != nil {
		return uuid.Nil, err
	}
	if r.TX != nil {
		prepare, err := r.TX.Prepare("UPDATE users SET login=$1, email=$2, password_hash=$3 WHERE id=$4")
		if err != nil {
			return uuid.Nil, err
		}
		_, err = prepare.Exec(user.Login, user.Email, user.PasswordHash, userUIID)
		if err != nil {
			return uuid.Nil, err
		}
		return user.ID, nil
	}
	prepare, err := r.DB.Prepare("UPDATE users SET login=$1, email=$2, password_hash=$3 WHERE id=$4")
	if err != nil {
		return uuid.Nil, err
	}
	_, err = prepare.Exec(user.Login, user.Email, user.PasswordHash, userUIID)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}
func (r *UserRepo) DeleteUser(id uuid.UUID) (uuid.UUID, error) {
	userUIID, err := id.MarshalBinary()
	if err != nil {
		return uuid.Nil, errors.New("user not found")
	}
	if r.TX != nil {
		_, err = r.TX.Exec("DELETE FROM users WHERE id=$1", userUIID)
		if err != nil {
			return uuid.Nil, err
		}
		return id, err
	}
	_, err = r.TX.Exec("DELETE FROM users WHERE id=$1", userUIID)
	if err != nil {
		return uuid.Nil, err
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
