package services

import "online-shop/internal/store"

type ProductWebService struct {
	store *store.Store
}

func NewProductWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
