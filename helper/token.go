package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims CustomClaims) (string, error)
	VerifyJWT(tokenString string) (*jwt.Token, error)
}

type TokenUseCaseImpl struct{}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims CustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodeToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return encodeToken, nil
}

func (t *TokenUseCaseImpl) VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
