package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"online-shop/internal/models"
)

type ProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) CreateCategory(category string) (int, error) {
	if category == "" {
		return 0, errors.New("incorrect category")
	}
	var categoryID int
	if p.TX != nil {
		prepare, err := p.TX.Prepare("INSERT INTO products_categories(category_name) VALUES ($1) RETURNING id " +
			"WHERE NOT EXISTS (SELECT category_name FROM products_categories category_name=($1))")
		if err != nil {
			return 0, err
		}
		err = prepare.QueryRow(category).Scan(&categoryID)
		if err != nil {
			return 0, err
		}
		return categoryID, nil
	}
	prepare, err := p.DB.Prepare("INSERT INTO products_categories(category_name) VALUES ($1) RETURNING id " +
		"WHERE NOT EXISTS (SELECT category_name FROM products_categories category_name=($1))")
	if err != nil {
		return 0, err
	}
	err = prepare.QueryRow(category).Scan(&categoryID)
	if err != nil {
		return 0, err
	}
	return categoryID, nil

}

func (p *ProductRepo) CreateProduct(product *models.Product) (*models.Product, error) {
	if product == nil {
		return nil, errors.New("incorrect product")
	}
	if p.TX != nil {
		prepare, err := p.TX.Prepare("INSERT INTO products(product_name, external_product_id, product_category)" +
			"VALUES ($1, $2, (SELECT pc.id FROM products_categories AS pc WHERE pc.category_name=$3))")
		if err != nil {
			return nil, err
		}
		err = prepare.QueryRow(product.Name, product.ExternalProductID, product.Category).Scan(&product.ID)
		if err != nil {
			return nil, err
		}

		return product, err
	}
	prepare, err := p.DB.Prepare("INSERT INTO products(product_name, external_product_id, product_category)" +
		"VALUES ($1, $2, (SELECT pc.id FROM products_categories AS pc WHERE pc.category_name=$3)) RETURNING id")
	if err != nil {
		return nil, err
	}
	err = prepare.QueryRow(product.Name, product.ExternalProductID, product.Category).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, err
}
func (p *ProductRepo) CreateProductSupplier(product *models.Product) error {
	if product == nil {
		return errors.New("incorrect menu")
	}
	if p.TX != nil {
		prepare, err := p.TX.Prepare("INSERT INTO products_suppliers(product_id, supplier_id, " +
			"external_product_id, external_supplier_id, price, image, description) " +
			"VALUES ((SELECT p.id FROM products AS p WHERE p.product_name=$1), " +
			"(SELECT s.id FROM suppliers AS s WHERE s.id=$2), $3,$4,$5,$6,$7)")
		if err != nil {
			return err
		}
		_, err = prepare.Exec(product.Name, product.SupplierID, product.ExternalProductID, product.ExternalSupplierID,
			product.Price, product.Image, product.Description)
		if err != nil {
			return err
		}
		return nil
	}
	prepare, err := p.DB.Prepare("INSERT INTO products_suppliers(product_id, supplier_id, " +
		"external_product_id, external_supplier_id, price, image, description) " +
		"VALUES ((SELECT p.id FROM products AS p WHERE p.product_name=$1), " +
		"(SELECT s.id FROM suppliers AS s WHERE s.id=$2), $3,$4,$5,$6,$7)")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(product.Name, product.SupplierID, product.ExternalProductID, product.ExternalSupplierID,
		product.Price, product.Image, product.Description)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepo) GetProductInfo(id int) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("incorrect data")
	}
	var product models.Product
	if p.TX != nil {
		prepare, err := p.TX.Prepare("SELECT SELECT p.id, p.product_name, pc.category_name, " +
			"s.external_product_id, s.external_supplier_id, s.price, s.image, s.description, s.quantity " +
			"FROM products as p LEFT JOIN products_categories pc on pc.id = p.product_category " +
			"LEFT JOIN products_suppliers s s.product_id=$1")
		if err != nil {
			return nil, err
		}
		err = prepare.QueryRow(id).Scan(&product.ID, &product.Name, &product.Category, &product.ExternalProductID,
			&product.ExternalSupplierID, &product.Price, &product.Image, &product.Description, &product.Quantity)
		if err != nil {
			return nil, err
		}
		return &product, nil

	}
	prepare, err := p.DB.Prepare("SELECT SELECT p.id, p.product_name, pc.category_name, " +
		"s.external_product_id, s.external_supplier_id, s.price, s.image, s.description, s.quantity " +
		"FROM products as p LEFT JOIN products_categories pc on pc.id = p.product_category " +
		"LEFT JOIN products_suppliers s on s.product_id=$1")
	if err != nil {
		return nil, err
	}
	err = prepare.QueryRow(id).Scan(&product.ID, &product.Name, &product.Category, &product.ExternalProductID,
		&product.ExternalSupplierID, &product.Price, &product.Image, &product.Description, &product.Quantity)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *ProductRepo) GetAllProducts() (*[]models.Product, error) {
	var productList []models.Product

	if p.TX != nil {
		rows, err := p.TX.Query("SELECT SELECT p.id, p.product_name, pc.category_name, " +
			"s.external_product_id, s.external_supplier_id, s.price, s.image, s.description, s.quantity " +
			"FROM products as p LEFT JOIN products_categories pc on pc.id = p.product_category " +
			"LEFT JOIN products_suppliers s on p.id = s.product_id")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product models.Product
			err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.ExternalProductID,
				&product.ExternalSupplierID, &product.Price, &product.Image, &product.Description, &product.Quantity)

			if err != nil {
				return nil, err
			}
			productList = append(productList, product)
		}
		return &productList, nil

	}
	rows, err := p.DB.Query("SELECT SELECT p.id, p.product_name, pc.category_name, " +
		"s.external_product_id, s.external_supplier_id, s.price, s.image, s.description, s.quantity " +
		"FROM products as p LEFT JOIN products_categories pc on pc.id = p.product_category " +
		"LEFT JOIN products_suppliers s on p.id = s.product_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.ExternalProductID,
			&product.ExternalSupplierID, &product.Price, &product.Image, &product.Description, &product.Quantity)

		if err != nil {
			return nil, err
		}
		productList = append(productList, product)
	}
	return &productList, nil

}

func (p *ProductRepo) UpdatePrice(id int, price float32) error {
	if id <= 0 || price <= 0 {
		str := fmt.Sprintf("invalid data, %d, %f", id, price)
		return errors.New(str)
	}
	if p.TX != nil {
		_, err := p.TX.Exec("UPDATE products_suppliers SET price= $2 FROM products_suppliers "+
			"INNER JOIN products p ON p.id = $1", id, price)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := p.DB.Exec("UPDATE products_suppliers SET price= $2 FROM products_suppliers "+
		"INNER JOIN products p ON p.id = $1", id, price)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepo) UpdateQuantity(id int, quantity float32) error {
	if id <= 0 || quantity <= 0 {
		str := fmt.Sprintf("invalid data, %d, %f", id, quantity)
		return errors.New(str)
	}
	if p.TX != nil {
		_, err := p.TX.Exec("UPDATE products_suppliers SET quantity= $2 FROM products_suppliers "+
			"INNER JOIN products p ON p.id = $1", id, quantity)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := p.DB.Exec("UPDATE products_suppliers SET quantity= $2 FROM products_suppliers "+
		"INNER JOIN products p ON p.id = $1", id, quantity)
	if err != nil {
		return err
	}
	return nil
}

//todo
//func (p *ProductRepo) DeleteProduct(id int) error                    {
//	if id <= 0{
//		str := fmt.Sprintf("invalid id, %d", id)
//		return errors.New(str)
//	}
//	return nil
//}

func (p *ProductRepo) BeginTX() error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	p.TX = tx
	return nil
}
func (p *ProductRepo) RollbackTX() error {
	defer func() {
		p.TX = nil
	}()
	if p.TX != nil {
		return p.TX.Rollback()
	}
	return nil
}
func (p *ProductRepo) CommitTX() error {
	defer func() {
		p.TX = nil
	}()
	if p.TX != nil {
		return p.TX.Commit()
	}
	return nil
}
func (p *ProductRepo) SetTX(tx *sql.Tx) {
	p.TX = tx
}
func (p *ProductRepo) GetTX() *sql.Tx {
	return p.TX
}
