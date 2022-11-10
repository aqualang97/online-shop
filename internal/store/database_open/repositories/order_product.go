package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"online-shop/internal/models"
)

type OrderProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderProductRepo(db *sql.DB) *OrderProductRepo {
	return &OrderProductRepo{DB: db}
}

// Переделать!!!!
func (o *OrderProductRepo) AddProduct(productID, orderID uuid.UUID, numbers int) error {
	if productID == uuid.Nil || orderID == uuid.Nil || numbers <= 0 {
		return errors.New("invalid arg(s)")
	}
	var inStock int
	pid, err := productID.MarshalBinary()
	if err != nil {
		return err
	}
	oid, err := orderID.MarshalBinary()
	if err != nil {
		return err
	}
	var price float32
	if o.TX != nil {
		prepare, err := o.TX.Prepare("select quantity, price from products_suppliers where product_id=$1")
		if err != nil {
			return err
		}
		err = prepare.QueryRow(pid).Scan(&inStock, &price)
		if err != nil {
			return err
		}
		if inStock < numbers {
			str := fmt.Sprintf("in stock only %d", inStock)
			return errors.New(str)
		}
		_, err = o.TX.Exec(
			"INSERT INTO orders_products(product_id, order_id, numbers_of_products, purchase_price) "+
				"VALUES ($1,$2,$3,$4 (SELECT price FROM products_suppliers WHERE product_id=$1))", pid, oid, numbers, float32(numbers)*price)
		if err != nil {
			return err
		}
		return nil
	}
	prepare, err := o.DB.Prepare("select quantity, price from products_suppliers where product_id=$1")
	if err != nil {
		return err
	}
	err = prepare.QueryRow(pid).Scan(&inStock, &price)
	if err != nil {
		return err
	}
	if inStock < numbers {
		str := fmt.Sprintf("in stock only %d", inStock)
		return errors.New(str)
	}
	_, err = o.DB.Exec(
		"INSERT INTO orders_products(product_id, order_id, numbers_of_products, purchase_price) "+
			"VALUES ($1,$2,$3, (SELECT price FROM products_suppliers WHERE product_id=$1))", pid, oid, numbers, float32(numbers)*price)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderProductRepo) UpdateProduct(productID, orderID uuid.UUID, numbers int) error {
	if productID == uuid.Nil || orderID == uuid.Nil || numbers <= 0 {
		return errors.New("invalid arg(s)")
	}
	var inStock, inOrder int
	pid, err := productID.MarshalBinary()
	if err != nil {
		return err
	}
	oid, err := orderID.MarshalBinary()
	if err != nil {
		return err
	}
	//var price float32

	if o.TX != nil {
		prepare, err := o.TX.Prepare("select ps.quantity, op.numbers_of_products, ps.price from products_suppliers as ps " +
			"left join orders_products op on ps.product_id = $1 and op.order_id=1")
		if err != nil {
			return err
		}
		err = prepare.QueryRow(pid, oid).Scan(&inStock, &inOrder)
		if err != nil {
			return err
		}

		if inStock+inOrder < numbers {
			str := fmt.Sprintf("in stock only %d", inStock)
			return errors.New(str)
		}
		_, err = o.TX.Exec(
			"INSERT INTO orders_products(product_id, order_id, numbers_of_products, purchase_price) " +
				"VALUES ($1,$2,$3, (SELECT price FROM products_suppliers WHERE product_id=$1))")
		if err != nil {
			return err
		}
		return nil
	}
	prepare, err := o.TX.Prepare("select ps.quantity, op.numbers_of_products from products_suppliers as ps " +
		"left join orders_products op on ps.product_id = $1 and op.order_id=1")
	if err != nil {
		return err
	}
	err = prepare.QueryRow(productID, orderID).Scan(&inStock, &inOrder)
	if err != nil {
		return err
	}

	if inStock+inOrder < numbers {
		str := fmt.Sprintf("in stock only %d", inStock)
		return errors.New(str)
	}
	_, err = o.TX.Exec(
		"INSERT INTO orders_products(product_id, order_id, numbers_of_products, purchase_price) " +
			"VALUES ($1,$2,$3, (SELECT price FROM products_suppliers WHERE product_id=$1))")
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderProductRepo) GetAllProduct(orderID int) (*[]models.OrderProduct, error) {
	var orderList []models.OrderProduct
	if o.TX != nil {
		rows, err := o.TX.Query("SELECT product_id, order_id, numbers_of_products, purchaise_prise")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var order models.OrderProduct
			rows.Scan(&order.ProductID, &order.OrderID, &order.NumbersOfProducts, &order.PurchasePrice)
			orderList = append(orderList, order)
		}
		return &orderList, nil
	}
	rows, err := o.DB.Query("SELECT product_id, order_id, numbers_of_products, purchaise_prise")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order models.OrderProduct
		rows.Scan(&order.ProductID, &order.OrderID, &order.NumbersOfProducts, &order.PurchasePrice)
		orderList = append(orderList, order)
	}
	return &orderList, nil
}
func (o *OrderProductRepo) DeleteProduct(productID, orderID int) error {
	if o.TX != nil {
		_, err := o.TX.Exec("DELETE FROM orders_products WHERE product_id=$1 AND order_id=$2", productID, orderID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := o.DB.Exec("DELETE FROM orders_products WHERE product_id=$1 AND order_id=$2", productID, orderID)
	if err != nil {
		return err
	}
	return nil

}
func (o *OrderProductRepo) DeleteAllProducts(orderID int) error {
	if o.TX != nil {
		_, err := o.TX.Exec("DELETE * FROM orders_products WHERE order_id=$1", orderID)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := o.DB.Exec("DELETE FROM orders_products WHERE order_id=$1", orderID)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderProductRepo) BeginTX() error {
	tx, err := o.DB.Begin()
	if err != nil {
		return err
	}
	o.TX = tx
	return nil
}
func (o *OrderProductRepo) RollbackTX() error {
	defer func() {
		o.TX = nil
	}()
	if o.TX != nil {
		return o.TX.Rollback()
	}
	return nil
}
func (o *OrderProductRepo) CommitTX() error {
	defer func() {
		o.TX = nil
	}()
	if o.TX != nil {
		return o.TX.Commit()
	}
	return nil
}
func (o *OrderProductRepo) SetTX(tx *sql.Tx) {
	o.TX = tx
}
func (o *OrderProductRepo) GetTX() *sql.Tx {
	return o.TX
}
