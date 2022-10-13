package models

import (
	"time"
)

type UserData struct {
	UserID   int    `json:"userID"`
	FullName string `json:"fullName"`
	//to change
	DateOfBirth     string     `json:"dateOfBirth"`
	Number          string     `json:"number"`
	Address         string     `json:"address"`
	DiscountPercent int        `json:"discountPercent"`
	DiscountAmount  float32    `json:"discountAmount"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
