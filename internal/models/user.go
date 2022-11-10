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
	Role         string     `json:"role"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}

type UserRegistrationRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserByRegData(urr *UserRegistrationRequest) (*User, error) {

	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(urr.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           userID,
		Login:        urr.Login,
		Email:        urr.Email,
		PasswordHash: string(passwordHashByte),
	}, nil
}
