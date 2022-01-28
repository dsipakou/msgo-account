package utils

import (
	"log"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("3djIDfjer454DFe3fdc")

func GetHash(pwd string) string {
  password := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
