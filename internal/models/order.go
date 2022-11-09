package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID              uuid.UUID  `json:"ID"`
	UserID          uuid.UUID  `json:"userID"`
	TotalPrice      float32    `json:"totalPrice"`
	Status          string     `json:"status"`
	PaymentMethod   string     `json:"paymentMethod"`
	DiscountAmount  float32    `json:"discountAmount"`
	DiscountPercent int        `json:"discountPercent"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}
