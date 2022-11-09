package services

import (
	"errors"
	"github.com/google/uuid"
	"online-shop/internal/models"
	"online-shop/internal/store"
)

type UsersWebService struct {
	store *store.Store
}

func NewUserWebService(store *store.Store) *UsersWebService {
	return &UsersWebService{
		store: store,
	}
}
func (u *UsersWebService) CreateUser(user *models.User, token *models.Token) (uuid.UUID, error) {
	err := u.store.Users.BeginTX()
	if err != nil {
		return uuid.Nil, err
	}
	u.store.Tokens.SetTX(u.store.Users.GetTX())
	userID, err := u.store.Users.CreateUser(user)
	if err != nil {
		_ = u.store.Users.RollbackTX()
		return uuid.Nil, errors.New("this email is already taken")
	}
	token.UserID = userID
	err = u.store.Tokens.CreateToken(token)
	if err != nil {
		_ = u.store.Tokens.RollbackTX()
		return uuid.Nil, err
	}
	err = u.store.Users.CommitTX()
	if err != nil {
		_ = u.store.Users.RollbackTX()
		return uuid.Nil, err
	}
	u.store.Users.SetTX(nil)
	u.store.Tokens.SetTX(nil)
	return userID, nil
}

func (u *UsersWebService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.store.Users.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil

}
