package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID                 uuid.UUID  `json:"ID"`
	SupplierID         uuid.UUID  `json:"supplierID"`
	ExternalProductID  int        `json:"externalProductID"`
	ExternalSupplierID int        `json:"externalSupplierID"`
	Name               string     `json:"name"`
	Category           string     `json:"category"`
	Price              float32    `json:"price"`
	Image              string     `json:"image"`
	Description        string     `json:"description"`
	Quantity           int        `json:"quantity"`
	CreatedAt          *time.Time `json:"createdAt"`
}
