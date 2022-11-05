package models

import (
	"time"
)

type Order struct {
	ID              int        `json:"ID"`
	UserID          int        `json:"userID"`
	TotalPrice      float32    `json:"totalPrice"`
	Status          string     `json:"status"`
	PaymentMethod   string     `json:"paymentMethod"`
	DiscountAmount  float32    `json:"discountAmount"`
	DiscountPercent int        `json:"discountPercent"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
