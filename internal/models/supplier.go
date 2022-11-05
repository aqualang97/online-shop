package models

import (
	"time"
)

type Supplier struct {
	ID                 int
	ExternalSupplierID int
	Image              string
	Description        string
	SupplierName       string
	CreatedAt          *time.Time
}
