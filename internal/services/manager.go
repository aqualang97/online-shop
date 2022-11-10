package services

import (
	"errors"
	"online-shop/internal/store"
)

type Manager struct {
	Admin        AdminService
	Order        OrderService
	OrderProduct OrderProductService
	Product      ProductService
	Supplier     SupplierService
	Token        TokenService
	User         UserService
	UserData     UserDataService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &Manager{
		Admin:        NewAdminWebService(store),
		Order:        NewOrderWebService(store),
		OrderProduct: NewOrderProductWebService(store),
		Product:      NewProductWebService(store),
		Supplier:     NewSupplierWebService(store),
		Token:        NewTokenWebService(store),
		User:         NewUserWebService(store),
		UserData:     NewUserDataWebService(store),
	}, nil
}
