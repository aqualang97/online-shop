package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type JWTCustomClaims struct {
	UserID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID uuid.UUID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func GetHashOfToken(tokenString string) (string, error) {
	hashToken, err := bcrypt.GenerateFromPassword([]byte(tokenString), bcrypt.DefaultCost)
	if err != nil {
		log.Println()
	}
	return string(hashToken), err
}

func CompareHashTokenDBAndRequest(hashTokenDB, tokenReq string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashTokenDB), []byte(tokenReq)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetTokenFromBearerString(bearerString string) string {
	log.Println(bearerString)
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer ")
	if len(parts) != 2 {
		return ""
	}
	token := strings.TrimSpace(parts[1])

	if len(token) < 1 {
		return ""
	}

	return token
}

func ValidateToken(tokenString, secret string) (bool, error) {

	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error,
		) {
			return []byte(secret), nil
		})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("Invalid token")
	}

	return true, nil
}

func Claims(tokenString, secret string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwt.Token) (interface{}, error,
		) {
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}
	//without valid check
	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, errors.New("failed to parse token claims")
	}
	return claims, nil
}
