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

type Authenticator interface {
	AuthByIp(ctx context.Context, ip string) (*Account, error)
	AuthByPassword(ctx context.Context, login string, password string) (*Account, error)
	GetByLogin(login string) (*Account, error)
}
