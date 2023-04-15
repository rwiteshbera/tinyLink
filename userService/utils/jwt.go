package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(email string, jwtSecret string) (signedToken string, err error) {
	claims := SignedDetails{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)), // Token will be expired after 168 H
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	return token, err
}

// Validate JWT Token
func ValidateToken(signedToken string, jwtSecret string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		msg = err.Error()
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = "token is invalid"
	}

	return claims, msg
}
