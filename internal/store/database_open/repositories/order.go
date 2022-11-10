package repositories

import (
	"database/sql"
)

type OrderRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

const empty = "empty"
const awaitingPayment = "awaiting payment"
const paid = "paid"
const completed = "completed"

//
//func (o *OrderRepo) NewEmptyOrder(userID uuid.UUID) (uuid.UUID, error) {
//	uid, err := userID.MarshalBinary()
//	if err != nil {
//		return uuid.Nil, err
//	}
//	var orderID uuid.UUID
//	orderID = uuid.New()
//	if o.TX != nil {
//		prepare, err := o.TX.Prepare("INSERT INTO orders(user_id) VALUES ($1) RETURNING id")
//		if err != nil {
//			return 0, err
//		}
//		err = prepare.QueryRow(userID).Scan(&orderID)
//		if err != nil {
//			return 0, err
//		}
//		return orderID, nil
//	}
//	prepare, err := o.DB.Prepare("INSERT INTO orders(user_id) VALUES ($1) RETURNING id")
//	if err != nil {
//		return 0, err
//	}
//	err = prepare.QueryRow(userID).Scan(&orderID)
//	if err != nil {
//		return 0, err
//	}
//	return orderID, nil
//}

//todo
//func (o *OrderRepo) SetStatus(orderID int, status string) (*models.Order, error) {
//	switch status {
//	case empty:
//		return &models.Order{}, nil
//	case awaitingPayment:
//	case paid:
//	case completed:
//	default:
//		return nil, errors.New("invalid data")
//	}
//
//}

func (o *OrderRepo) BeginTX() error {
	tx, err := o.DB.Begin()
	if err != nil {
		return err
	}
	o.TX = tx
	return nil
}
func (o *OrderRepo) RollbackTX() error {
	defer func() {
		o.TX = nil
	}()
	if o.TX != nil {
		return o.TX.Rollback()
	}
	return nil
}
func (o *OrderRepo) CommitTX() error {
	defer func() {
		o.TX = nil
	}()
	if o.TX != nil {
		return o.TX.Commit()
	}
	return nil
}
func (o *OrderRepo) SetTX(tx *sql.Tx) {
	o.TX = tx
}
func (o *OrderRepo) GetTX() *sql.Tx {
	return o.TX
}
