package app

import (
	"golang.org/x/crypto/bcrypt"
)

type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

func (hs *HashService) HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (hs *HashService) ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
