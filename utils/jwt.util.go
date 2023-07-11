package utils

import (
	"github.com/golang-jwt/jwt"
)

var SecretKey = "SECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	tokenChan := make(chan string)
	errChan := make(chan error)

	go func() {
		// Create Handler
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		webtoken, err := token.SignedString([]byte(SecretKey))
		if err != nil {
			errChan <- err
			return
		}
		tokenChan <- webtoken
	}()

	select {
	case token := <-tokenChan:
		return token, nil
	case err := <-errChan:
		return "", err
	}
}
