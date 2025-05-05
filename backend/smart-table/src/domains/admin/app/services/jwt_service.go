package app

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/smart-table/src/dependencies"
	appErrors "github.com/smart-table/src/domains/admin/app/services/errors"
	"github.com/smart-table/src/logging"
)

type JwtService struct {
	secretKey       []byte
	tokenExpiration time.Duration
}

func NewJwtService(deps *dependencies.Dependencies) *JwtService {
	return &JwtService{
		secretKey:       []byte(deps.Config.App.Admin.Jwt.SecretKey),
		tokenExpiration: deps.Config.App.Admin.Jwt.Expiration,
	}
}

type UserClaims struct {
	UserUUID uuid.UUID `json:"user_uuid"`
	jwt.RegisteredClaims
}

func (js *JwtService) GenerateJWT(userUUID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(js.tokenExpiration)

	claims := &UserClaims{
		UserUUID: userUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(js.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (js *JwtService) ValidateJWT(tokenString string, userUUID uuid.UUID) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return js.secretKey, nil
	})

	logger := logging.GetLogger()

	switch {
	case err != nil:
		return nil, err
	case claims.UserUUID != userUUID:
		logger.Error(fmt.Sprintf("Token.user_uuid mismatch with Header.user_uuid=%v", userUUID))
		return nil, appErrors.InvalidToken{}
	case !token.Valid:
		logger.Error("Invalid token")
		return nil, appErrors.InvalidToken{}
	}

	return claims, nil
}
