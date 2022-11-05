package models

import (
	"time"
)

type Token struct {
	ID               int        `json:"ID"`
	UserID           int        `json:"userID"`
	AccessTokenHash  string     `json:"accessTokenHash"`
	RefreshTokenHash string     `json:"refreshTokenHash"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"createdAt"`
}
