package models

import (
	"golang.org/x/crypto/bcrypt"

	"time"
)

type User struct {
	ID           int        `json:"ID"`
	Login        string     `json:"login"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"passwordHash"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

func CreateUserByRegData(login, email, password string) (*User, error) {
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Login:        login,
		Email:        email,
		PasswordHash: string(passwordHashByte),
	}, nil
}
