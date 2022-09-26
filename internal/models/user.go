package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"time"
)

type User struct {
	ID           uuid.UUID  `json:"ID"`
	Login        string     `json:"login"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"passwordHash"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

func CreateUserByRegData(login, email, password string) (*User, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           userUUID,
		Login:        login,
		Email:        email,
		PasswordHash: string(passwordHashByte),
	}, nil
}
