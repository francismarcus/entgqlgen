package auth

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates a hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CompareHash compares the hash to the password
func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken generates JWT
func GenerateToken(id int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return t, nil
}

// DecodeToken decodes the token and returns the claims
func DecodeToken(t string) (jwt.MapClaims, bool) {
	str := os.Getenv("JWT_SECRET")
	secret := []byte(str)

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return secret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	log.Printf("Invalid JWT Token")
	return nil, false
}
