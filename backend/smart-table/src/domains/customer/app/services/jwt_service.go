package app

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/smart-table/src/dependencies"
	appErrors "github.com/smart-table/src/domains/customer/app/services/errors"
	"github.com/smart-table/src/logging"
)

type JwtService struct {
	secretKey       []byte
	tokenExpiration time.Duration
}

func NewJwtService(deps *dependencies.Dependencies) *JwtService {
	return &JwtService{
		secretKey:       []byte(deps.Config.App.Customer.Jwt.SecretKey),
		tokenExpiration: deps.Config.App.Customer.Jwt.Expiration,
	}
}

type CustomerClaims struct {
	CustomerUUID uuid.UUID `json:"customer_uuid"`
	jwt.RegisteredClaims
}

func (js *JwtService) GenerateJWT(customerUUID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(js.tokenExpiration)

	claims := &CustomerClaims{
		CustomerUUID: customerUUID,
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

func (js *JwtService) ValidateJWT(tokenString string, customerUUID uuid.UUID) error {
	claims := &CustomerClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return js.secretKey, nil
	})

	logger := logging.GetLogger()

	switch {
	case err != nil:
		return err
	case claims.CustomerUUID != customerUUID:
		logger.Error(fmt.Sprintf("Token.customer_uuid=%v mismatch with Header.customer_uuid=%v", claims.CustomerUUID, customerUUID))
		return appErrors.InvalidToken{}
	case !token.Valid:
		logger.Error("Invalid token")
		return appErrors.InvalidToken{}
	}

	return nil
}
