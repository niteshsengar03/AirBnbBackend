package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(PlainPassword string) (string, error) {
	pass := []byte(PlainPassword)
	hass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing the password")
		return "", err
	}
	return string(hass), nil
}

func CheckpasswordHash(PlainPassword string, HashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(PlainPassword))
	return err == nil
}
