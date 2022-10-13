package models

import (
	"time"
)

type OrderProduct struct {
	ID                int
	ProductID         int
	OrderID           int
	NumbersOfProducts int
	PurchasePrice     float32
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
