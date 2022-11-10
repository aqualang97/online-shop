package services

import (
	"online-shop/internal/store"
)

type AdminWebService struct {
	store *store.Store
}

func NewAdminWebService(store *store.Store) *AdminWebService {
	return &AdminWebService{
		store: store,
	}
}

func (a *AdminWebService) IsAdmin(login string) (bool, error) {
	admin, err := a.store.Admin.CheckRole(login)
	if err != nil {
		return false, err
	}

	return admin, nil
}
