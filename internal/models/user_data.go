package models

import (
	"github.com/google/uuid"
	"time"
)

type UserData struct {
	UserID   uuid.UUID `json:"userID"`
	FullName string    `json:"fullName"`
	//to change
	DateOfBirth     string     `json:"dateOfBirth"`
	Number          string     `json:"number"`
	Address         string     `json:"address"`
	DiscountPercent int        `json:"discountPercent"`
	DiscountAmount  float32    `json:"discountAmount"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
