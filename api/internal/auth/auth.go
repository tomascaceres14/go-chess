package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessCredentials struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// Interfaced for future integrations with OAuth / third-party auth provider
type TokenProvider interface {
	NewUserCredentials(id string) (*AccessCredentials, error)
	ValidateToken(token string) (*jwt.Token, error)
}

// Custom JWT implementation
type JWTTokenProvider struct {
	signingKey []byte
	iss        string
}

func NewJWTTokenProvider(signingKey string, iss string) (*JWTTokenProvider, error) {

	key := []byte(signingKey)
	return &JWTTokenProvider{
		signingKey: key,
		iss:        iss,
	}, nil
}

func (j *JWTTokenProvider) NewUserCredentials(id string) (*AccessCredentials, error) {
	at, err := j.newAccessToken(id)
	if err != nil {
		return nil, err
	}

	rt, err := j.newRefreshToken(id)
	if err != nil {
		return nil, err
	}

	return &AccessCredentials{
		Token:        at,
		RefreshToken: rt,
	}, nil
}

func (j *JWTTokenProvider) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token,
		&jwt.MapClaims{},
		func(t *jwt.Token) (any, error) {
			return j.signingKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}), jwt.WithIssuer(j.iss), jwt.WithExpirationRequired(),
	)
}

func (j *JWTTokenProvider) newAccessToken(id string) (string, error) {
	return j.newToken(id, map[string]any{
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"type": "access",
	})
}

func (j *JWTTokenProvider) newRefreshToken(id string) (string, error) {
	return j.newToken(id, map[string]any{
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"type": "refresh",
	})
}

func (j *JWTTokenProvider) newToken(id string, claims jwt.MapClaims) (string, error) {

	// Default claims
	claims["iss"] = j.iss
	claims["sub"] = id
	claims["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		claims,
	)

	s, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", err
	}

	return s, nil
}
