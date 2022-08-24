package models

import (
	"github.com/google/uuid"
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
