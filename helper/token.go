package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func CreateJWTToken(username string, isAdmin bool, secret string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &JwtClaims{
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Can't sign token: ", err.Error())
		return "", err
	}

	return tokenString, nil

}
