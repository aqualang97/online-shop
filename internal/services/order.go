package services

import "online-shop/internal/store"

type OrderWebService struct {
	store *store.Store
}

func NewOrderWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
