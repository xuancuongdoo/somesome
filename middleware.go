package main

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt"
)

func withJWTAuth(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Checking JWT token")
		tokenString := request.Header.Get("x-jwt-token")
		_, err := validateJWT(tokenString)
		if err != nil {
			WriteJSON(response, http.StatusUnauthorized, ApiError{Error: "Invalid Token"})
			return
		}
		handleFunc(response, request)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}


func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"accountNumber" : account.Number,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}