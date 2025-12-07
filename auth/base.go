package auth

import (
	"net/http"
)

type Account struct {
	Login       string
	ProfileSlug string
}

type HttpAuthenticator interface {
	GetUser(r *http.Request) (*Account, error)
}
