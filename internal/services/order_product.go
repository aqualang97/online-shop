package services

import "online-shop/internal/store"

type OrderProductWebService struct {
	store *store.Store
}

func NewOrderProductWebService(store *store.Store) *OrderProductWebService {
	return &OrderProductWebService{
		store: store,
	}
}
