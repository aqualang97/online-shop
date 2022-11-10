package store

import (
	"database/sql"
	"online-shop/internal/store/database_open/repositories"
)

type Store struct {
	Admin         *repositories.AdminRepo
	Orders        *repositories.OrderRepo
	OrderProducts *repositories.OrderProductRepo
	Products      *repositories.ProductRepo
	Suppliers     *repositories.SupplierRepo
	Tokens        *repositories.TokenRepo
	Users         *repositories.UserRepo
	UserData      *repositories.UserDataRepo
}

func NewStore(db *sql.DB) *Store {
	ar := repositories.NewAdminRepo(db)
	or := repositories.NewOrderRepo(db)
	opr := repositories.NewOrderProductRepo(db)
	pr := repositories.NewProductRepo(db)
	sr := repositories.NewSupplierRepo(db)
	tr := repositories.NewTokenRepo(db)
	ur := repositories.NewUserRepo(db)
	udr := repositories.NewUserDataRepo(db)

	return &Store{
		Admin:         ar,
		Orders:        or,
		OrderProducts: opr,
		Products:      pr,
		Suppliers:     sr,
		Tokens:        tr,
		Users:         ur,
		UserData:      udr,
	}
}
