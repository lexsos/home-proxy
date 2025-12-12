package auth

import (
	"context"
	"net/http"
)

type Account struct {
	Login       string
	ProfileSlug string
}

type HttpAuthenticator interface {
	GetUser(ctx context.Context, r *http.Request) (*Account, error)
}
