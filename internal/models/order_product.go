package models

import (
	"github.com/google/uuid"
	"time"
)

type OrderProduct struct {
	ID                uuid.UUID
	ProductID         uuid.UUID
	OrderID           uuid.UUID
	NumbersOfProducts int
	PurchasePrice     float32
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
