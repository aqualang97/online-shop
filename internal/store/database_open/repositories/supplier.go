package repositories

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"online-shop/internal/models"
)

type SupplierRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewSupplierRepo(db *sql.DB) *SupplierRepo {
	return &SupplierRepo{DB: db}
}

func (s *SupplierRepo) CreateSupplier(supp *models.Supplier) (uuid.UUID, error) {
	if supp == nil {
		return uuid.Nil, errors.New("invalid data")
	}
	sid := uuid.New()
	sidB, err := sid.MarshalBinary()
	if err != nil {
		return uuid.Nil, err
	}
	if s.TX != nil {
		prepare, err := s.TX.Prepare("INSERT INTO suppliers(id, external_supplier_id, supplier_name, image, description) " +
			"VALUES ($1,$2,$3,$4,$5)")
		if err != nil {
			return uuid.Nil, err
		}
		_, err = prepare.Exec(sidB, supp.ExternalSupplierID, supp.SupplierName, supp.Image, supp.Description)
		if err != nil {
			return uuid.Nil, err
		}
		return sid, nil
	}
	prepare, err := s.DB.Prepare("INSERT INTO suppliers(id, external_supplier_id, supplier_name, image, description) " +
		"VALUES ($1,$2,$3,$4, $5) RETURNING id")
	if err != nil {
		return uuid.Nil, err
	}
	_, err = prepare.Exec(sidB, supp.ExternalSupplierID, supp.SupplierName, supp.Image, supp.Description)
	if err != nil {
		return uuid.Nil, err
	}
	return sid, nil
}

//todo
//func (s *SupplierRepo) UpdateSupplier()  {}

func (s *SupplierRepo) GetSupplier(id uuid.UUID) (*models.Supplier, error) {
	var supp models.Supplier
	sid, err := id.MarshalBinary()
	if s.TX != nil {
		prepare, err := s.TX.Prepare("SELECT id, external_supplier_id, supplier_name, image, description FROM suppliers WHERE id=$1")
		if err != nil {
			return nil, err
		}
		err = prepare.QueryRow(sid).Scan(&supp.ID, &supp.ExternalSupplierID, &supp.SupplierName, &supp.Image, &supp.Description)
		if err != nil {
			return nil, err
		}
		return &supp, nil
	}
	prepare, err := s.DB.Prepare("SELECT id, external_supplier_id, supplier_name, image, description FROM suppliers WHERE id=$1")
	if err != nil {
		return nil, err
	}
	err = prepare.QueryRow(sid).Scan(&supp.ID, &supp.ExternalSupplierID, &supp.SupplierName, &supp.Image, &supp.Description)
	if err != nil {
		return nil, err
	}
	return &supp, nil

}
func (s *SupplierRepo) GetAllSupplies() (*[]models.Supplier, error) {
	var suppList []models.Supplier
	if s.TX != nil {
		rows, err := s.TX.Query("SELECT id, external_supplier_id, supplier_name, image, description FROM suppliers")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var supp models.Supplier
			rows.Scan(&supp.ID, &supp.ExternalSupplierID, &supp.SupplierName, &supp.Image, &supp.Description)
			suppList = append(suppList, supp)
		}
		return &suppList, nil
	}
	rows, err := s.DB.Query("SELECT id, external_supplier_id, supplier_name, image, description FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var supp models.Supplier
		rows.Scan(&supp.ID, &supp.ExternalSupplierID, &supp.SupplierName, &supp.Image, &supp.Description)
		suppList = append(suppList, supp)
	}
	return &suppList, nil
}

func (s *SupplierRepo) BeginTX() error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	s.TX = tx
	return nil
}
func (s *SupplierRepo) RollbackTX() error {
	defer func() {
		s.TX = nil
	}()
	if s.TX != nil {
		return s.TX.Rollback()
	}
	return nil
}
func (s *SupplierRepo) CommitTX() error {
	defer func() {
		s.TX = nil
	}()
	if s.TX != nil {
		return s.TX.Commit()
	}
	return nil
}
func (s *SupplierRepo) SetTX(tx *sql.Tx) {
	s.TX = tx
}
func (s *SupplierRepo) GetTX() *sql.Tx {
	return s.TX
}
