package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	NewToken(uid string)
}

type JWTAuth struct {
	signingKey string
}

func NewJWTAuth(sKey string) *JWTAuth {
	return &JWTAuth{
		signingKey: sKey,
	}
}

func (j *JWTAuth) NewToken(uid string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":     "go-chess",
			"user_id": uid,
		},
	)

	s, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", err
	}

	return s, nil
}
