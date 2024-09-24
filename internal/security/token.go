package security

import (
	"fmt"
	"time"
	"user-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"roel"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(claim TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		ID:       claim.ID,
		Username: claim.Username,
		Role:     claim.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
		},
	})

	return token.SignedString([]byte(config.Load().SECRET_KEY))
}

func ExtractClaims(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
