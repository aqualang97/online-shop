package models

import "time"

type Supplier struct {
	ID                 int
	ExternalSupplierID int
	SupplierName       string
	CreatedAt          *time.Time
}
