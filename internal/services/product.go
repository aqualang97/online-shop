package services

import "online-shop/internal/store"

type ProductWebService struct {
	store *store.Store
}

func NewProductWebService(store *store.Store) *ProductWebService {
	return &ProductWebService{
		store: store,
	}
}
