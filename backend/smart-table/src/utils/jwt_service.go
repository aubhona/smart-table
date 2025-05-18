package utils

import (
	"github.com/google/uuid"
)

type JwtService interface {
	GenerateJWT(userUUID uuid.UUID) (string, error)
	ValidateJWT(tokenString string, userUUID uuid.UUID) error
}
