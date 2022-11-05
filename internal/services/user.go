package services

import "online-shop/internal/store"

type UsersWebService struct {
	store *store.Store
}

func NewUserWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
