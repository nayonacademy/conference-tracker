package account

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//secret key
var secretKey = []byte("abcd1234!@#$")

// ArithmeticCustomClaims
type ArithmeticCustomClaims struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`

	jwt.StandardClaims
}

// jwtKeyFunc
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

// Sign
func Sign(name, uid string) (string, error) {
	expAt := time.Now().Add(time.Duration(60) * time.Minute).Unix()
	claims := ArithmeticCustomClaims{
		UserId: uid,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "system",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}