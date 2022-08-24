package models

import "time"

type Product struct {
	ID                int        `json:"ID"`
	ExternalProductID int        `json:"externalProductID"`
	SupplierID        int        `json:"supplierID"`
	Name              string     `json:"name"`
	Category          string     `json:"category"`
	Price             float32    `json:"price"`
	Image             string     `json:"image"`
	Description       string     `json:"description"`
	Quantity          int        `json:"quantity"`
	CreatedAt         *time.Time `json:"createdAt"`
}
