package security

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword parolni hash qiladi
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt error: %v", err)
	}

	return string(hashedPassword), nil
}

// CheckPasswordHash parolni tekshiradi
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Invalid password:", err)
		return false
	}
	return true
}
