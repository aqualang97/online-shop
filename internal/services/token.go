package services

import "online-shop/internal/store"

type TokenWebService struct {
	store *store.Store
}

func NewTokenWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
