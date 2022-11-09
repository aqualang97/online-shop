package models

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	//ID               int        `json:"ID"`
	UserID           uuid.UUID  `json:"userID"`
	AccessTokenHash  string     `json:"accessTokenHash"`
	RefreshTokenHash string     `json:"refreshTokenHash"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

type RespToken struct {
	UserID       uuid.UUID `json:"userID"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

type RequestToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
