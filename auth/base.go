package auth

import (
	"net/http"
)

type Account struct {
	Id    int64
	Login string
	Role  string
}

type HttpAuthenticator interface {
	GetUser(r *http.Request) (*Account, error)
}
