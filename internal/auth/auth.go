package auth

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Key int

const Ctxkey Key = 1

type Auth struct {
	privatekey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}
type TokenAuth interface {
	GenerateToken(claims jwt.RegisteredClaims) (string, error)
	ValidateToken(token string) (jwt.RegisteredClaims, error)
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (TokenAuth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private key or public key cannot be nil")
	}
	return &Auth{
		privatekey: privateKey,
		publicKey:  publicKey,
	}, nil
}
