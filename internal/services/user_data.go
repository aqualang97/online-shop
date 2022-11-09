package services

import "online-shop/internal/store"

type UsersDataWebService struct {
	store *store.Store
}

func NewUserDataWebService(store *store.Store) *UsersDataWebService {
	return &UsersDataWebService{
		store: store,
	}
}
