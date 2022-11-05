package services

import "online-shop/internal/store"

type SupplierWebService struct {
	store *store.Store
}

func NewSupplierWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
