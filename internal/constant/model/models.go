package model

import (
	"github.com/golang-jwt/jwt"
	errors "sms-gateway/internal/constant/rest/error_types"
	"time"
)

type SMS struct {
	To      []string `json:"to"`
	Content string   `json:"content"`
}

type OutGoingSMS struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type Send struct {
	Sms []SMS `json:"messages"`
}

type ClientLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Phone string
	jwt.StandardClaims
}

func (ts *CustomClaims) GenerateToken(issuer string) (*string, error) {
	claims := &CustomClaims{
		issuer,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secretKey"))
	if err != nil {

		return nil, errors.ErrGenerateTokenError.Wrap(err, errors.ErrorGenerateTokenError)
	}

	return &t, nil
}
