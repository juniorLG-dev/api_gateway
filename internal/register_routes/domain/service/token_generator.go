package service

import (
	"gateway/internal/configuration/handler_err"
	"gateway/internal/register_routes/domain/entities"

	"github.com/golang-jwt/jwt"

	"errors"
	"os"
	"strings"
	"time"
)

type TokenInfoDTO struct {
	ID          string
	ServiceName string
}

type TokenGenerator struct{}

func NewTokenGenerator() *TokenGenerator {
	return &TokenGenerator{}
}

func (tg *TokenGenerator) GenerateToken(apiService entities.APIService) (string, *handler_err.InfoErr) {
	secretKey := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"id":   apiService.GetID(),
		"name": apiService.GetName(),
		"exp":  time.Now().Add(time.Hour * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", &handler_err.InfoErr{
			Message: "error creating jwt",
			Err:     handler_err.ErrInternal,
		}
	}

	return tokenString, &handler_err.InfoErr{}
}

func (tg *TokenGenerator) CheckToken(tokenValue string) (TokenInfoDTO, *handler_err.InfoErr) {
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tg.removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}

		return nil, errors.New("invalid token")
	})
	if err != nil {
		return TokenInfoDTO{}, &handler_err.InfoErr{
			Message: err.Error(),
			Err:     handler_err.ErrInvalidInput,
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return TokenInfoDTO{}, &handler_err.InfoErr{
			Message: "invalid token",
			Err:     handler_err.ErrInvalidInput,
		}
	}

	return TokenInfoDTO{
		ID:          claims["id"].(string),
		ServiceName: claims["name"].(string),
	}, &handler_err.InfoErr{}
}

func (tg *TokenGenerator) removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
