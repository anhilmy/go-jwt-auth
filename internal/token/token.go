package token

import (
	"fmt"
	"go-jwt-auth/internal/config"
	"go-jwt-auth/internal/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(key int) (string, error)
	TokenValid(token string) (string, error)
}

type service struct {
	config config.Config
}

func NewService(config *config.Config) Service {
	return service{*config}
}

func (s service) GenerateToken(userId int) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.config.TokenHourLifeSpan))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	res := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := res.SignedString([]byte(s.config.JWTSigningKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s service) TokenValid(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWTSigningKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, err := claims.GetSubject()
		return sub, err
	} else {
		return "", errors.TokenInvalid{}
	}

}
