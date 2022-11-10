package services

import (
	"github.com/google/uuid"
	"online-shop/internal/models"
)

type AdminService interface {
	IsAdmin(login string) (bool, error)
}
type OrderService interface{}

type OrderProductService interface{}

type ProductService interface{}

type SupplierService interface{}

type TokenService interface {
	CreateToken(token *models.Token) error
	DeleteTokenByUserID(id uuid.UUID) error
	GetTokenByUserID(userID uuid.UUID) (*models.Token, error)
	UpdateToken(token *models.Token) error
}

type UserService interface {
	CreateUser(user *models.User, token *models.Token) (uuid.UUID, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}

type UserDataService interface{}
