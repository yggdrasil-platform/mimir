package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	GrantType string `json:"gty,omitempty"`
}

func CreateJWT(id string, sub string, iat time.Time, exp time.Time, gty string, sk string) (string, error) {
	if gty != "client_credentials" && gty != "password" {
		return "", fmt.Errorf("invalid grant type, expected 'client_credentials' or 'password', got: %s", gty)
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		GrantType: gty,
		StandardClaims: jwt.StandardClaims{
			Id: id,
			ExpiresAt: exp.Unix(),
			IssuedAt: iat.Unix(),
			Issuer: "mimir.yggdrasil",
			Subject: sub,
		},
	})

	return tkn.SignedString([]byte(sk))
}

func VerifyAndParseJWT(tkn string, sk string, clms *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(strings.TrimSpace(tkn), clms, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(sk), nil
	})
}
