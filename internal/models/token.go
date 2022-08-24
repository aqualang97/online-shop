package models

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID               uuid.UUID  `json:"ID"`
	UserID           uuid.UUID  `json:"userID"`
	AccessTokenHash  string     `json:"accessTokenHash"`
	RefreshTokenHash string     `json:"refreshTokenHash"`
	CreatedAt        *time.Time `json:"createdAt"`
}
