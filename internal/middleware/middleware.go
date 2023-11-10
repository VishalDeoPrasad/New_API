package middleware

import (
	"errors"
	"job-application-api/internal/auth"
)

type Mid struct {
	a auth.TokenAuth
}

func NewMid(a auth.TokenAuth) (Mid, error) {
	if a == nil {
		return Mid{}, errors.New("auth can't be nil")
	}
	return Mid{a: a}, nil
}
