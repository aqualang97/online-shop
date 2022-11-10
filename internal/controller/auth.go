package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"online-shop/config"
	"online-shop/internal/models"
	"online-shop/internal/services"
	"online-shop/pkg/token"
)

type AuthController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewAuthController(services *services.Manager, config *config.Config) *AuthController {
	return &AuthController{
		cfg:      config,
		services: services,
	}
}
func (a *AuthController) Login(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		req := models.UserLoginRequest{}
		if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		var user *models.User
		var err error

		if req.Login == "" {
			user, err = a.services.User.GetUserByEmail(req.Email)
			if err != nil {
				http.Error(ctx.Writer, "invalid credentials", http.StatusUnauthorized)
				return
			}
		} else {
			user, err = a.services.User.GetUserByLogin(req.Login)
			if err != nil {
				http.Error(ctx.Writer, "invalid credentials", http.StatusUnauthorized)
				return
			}
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(ctx.Writer, "invalid credentials", http.StatusUnauthorized)
			return
		}
		err = a.services.Token.DeleteTokenByUserID(user.ID)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		accessString, err := token.GenerateToken(user.ID, a.cfg.AccessTokenLifeTime, a.cfg.AccessSecret)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		//accessHash, err := token.GetHashOfToken(accessString)
		//if err != nil {
		//	http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
		//	return
		//}
		refreshString, err := token.GenerateToken(user.ID, a.cfg.RefreshTokenLifeTime, a.cfg.RefreshSecret)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		//refreshHash, err := token.GetHashOfToken(accessString)
		//if err != nil {
		//	http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
		//	return
		//}
		resp := &models.RespToken{
			UserID:       user.ID,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		toCreate := &models.Token{
			UserID:           user.ID,
			AccessTokenHash:  base64.StdEncoding.EncodeToString([]byte(accessString)),
			RefreshTokenHash: base64.StdEncoding.EncodeToString([]byte(refreshString)),
		}
		err = a.services.Token.CreateToken(toCreate)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}

		ctx.Writer.WriteHeader(http.StatusOK)

		json.NewEncoder(ctx.Writer).Encode(resp)
	default:
		http.Error(ctx.Writer, fmt.Sprintf("Only %s method", http.MethodPost), http.StatusMethodNotAllowed)
	}
}
func (a *AuthController) Logout(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		accessString := token.GetTokenFromBearerString(ctx.Request.Header.Get("Authorization"))
		claims, err := token.Claims(accessString, a.cfg.AccessSecret)

		if err != nil {
			fmt.Println(err)
			http.Error(ctx.Writer, err.Error(), http.StatusUnauthorized)
			return
		}
		err = a.services.Token.DeleteTokenByUserID(claims.UserID)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)
		json.NewEncoder(ctx.Writer).Encode("Successful Logout")

	default:
		http.Error(ctx.Writer, fmt.Sprintf("Only %s method", http.MethodPost), http.StatusMethodNotAllowed)
	}

}
func (a *AuthController) Registration(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		req := &models.UserRegistrationRequest{}

		if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := models.CreateUserByRegData(req)
		fmt.Println(user)
		accessString, err := token.GenerateToken(user.ID, a.cfg.AccessTokenLifeTime, a.cfg.AccessSecret)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		//accessHash, err := token.GetHashOfToken(accessString)
		//if err != nil {
		//	http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
		//	return
		//}
		refreshString, err := token.GenerateToken(user.ID, a.cfg.RefreshTokenLifeTime, a.cfg.RefreshSecret)
		if err != nil {
			http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
			return
		}
		//refreshHash, err := token.GetHashOfToken(accessString)
		//if err != nil {
		//	http.Error(ctx.Writer, "invalid operation", http.StatusInternalServerError)
		//	return
		//}
		resp := &models.RespToken{
			UserID:       user.ID,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		toCreate := &models.Token{
			UserID:           user.ID,
			AccessTokenHash:  base64.StdEncoding.EncodeToString([]byte(accessString)),
			RefreshTokenHash: base64.StdEncoding.EncodeToString([]byte(refreshString)),
		}
		_, err = a.services.User.CreateUser(user, toCreate)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx.Writer.WriteHeader(http.StatusOK)

		json.NewEncoder(ctx.Writer).Encode(resp)

	default:
		http.Error(ctx.Writer, fmt.Sprintf("Only %s method", http.MethodPost), http.StatusMethodNotAllowed)
	}
}
func (a *AuthController) Refresh(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:

		//req := new(models.RequestToken)
		//if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		//	http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		//	return
		//}
		//refreshTokenString := token.GetTokenFromBearerString(req.RefreshToken)
		refreshTokenString := token.GetTokenFromBearerString(ctx.GetHeader("Authorization"))

		_, err := token.ValidateToken(refreshTokenString, a.cfg.RefreshSecret)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, _ := token.Claims(refreshTokenString, a.cfg.RefreshSecret)
		tokenDB, err := a.services.Token.GetTokenByUserID(claims.UserID)
		if err != nil {
			http.Error(ctx.Writer, "invalid token", http.StatusUnauthorized)
			return
		}
		//if !token.CompareHashTokenDBAndRequest(tokenDB.RefreshTokenHash, refreshTokenString) {
		//	http.Error(ctx.Writer, "invalid token", http.StatusUnauthorized)
		//	return
		//}

		if tokenDB.RefreshTokenHash != base64.StdEncoding.EncodeToString([]byte(refreshTokenString)) {
			http.Error(ctx.Writer, "invalid token", http.StatusUnauthorized)
			return
		}
		accessString, err := token.GenerateToken(claims.UserID, a.cfg.AccessTokenLifeTime, a.cfg.AccessSecret)

		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		refreshString, err := token.GenerateToken(claims.UserID, a.cfg.RefreshTokenLifeTime, a.cfg.RefreshSecret)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		//accessHash, _ := token.GetHashOfToken(accessString)
		//refreshHash, _ := token.GetHashOfToken(refreshString)

		resp := &models.RespToken{
			UserID:       claims.UserID,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		toUpdate := &models.Token{
			UserID:           claims.UserID,
			AccessTokenHash:  base64.StdEncoding.EncodeToString([]byte(accessString)),
			RefreshTokenHash: base64.StdEncoding.EncodeToString([]byte(refreshString)),
		}
		err = a.services.Token.UpdateToken(toUpdate)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)

		json.NewEncoder(ctx.Writer).Encode(resp)
	default:
		http.Error(ctx.Writer, fmt.Sprintf("Only %s method", http.MethodPost), http.StatusMethodNotAllowed)
	}
}
