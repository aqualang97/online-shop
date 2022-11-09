package models

import (
	"github.com/google/uuid"
	"time"
)

type Supplier struct {
	ID                 uuid.UUID
	ExternalSupplierID int
	Image              string
	Description        string
	SupplierName       string
	CreatedAt          *time.Time
}
