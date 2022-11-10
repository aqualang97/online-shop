package services

import (
	"github.com/google/uuid"
	"online-shop/internal/models"
	"online-shop/internal/store"
)

type TokenWebService struct {
	store *store.Store
}

func NewTokenWebService(store *store.Store) *TokenWebService {
	return &TokenWebService{
		store: store,
	}
}

func (t *TokenWebService) CreateToken(token *models.Token) error {
	err := t.store.Tokens.CreateToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenWebService) GetTokenByUserID(userID uuid.UUID) (*models.Token, error) {
	token, err := t.store.Tokens.GetTokensByUserID(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (t *TokenWebService) UpdateToken(token *models.Token) error {
	err := t.store.Tokens.UpdateTokens(token)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenWebService) DeleteTokenByUserID(id uuid.UUID) error {
	err := t.store.Tokens.DeleteTokensByUserID(id)
	if err != nil {
		return err
	}
	return nil
}
