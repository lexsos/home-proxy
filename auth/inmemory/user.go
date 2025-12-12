package inmemory

import (
	"context"
	"net/http"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/logging"
	"github.com/lexsos/home-proxy/request"
)

func (jsonAuth *HttpAuthenticator) GetUser(ctx context.Context, r *http.Request) (*auth.Account, error) {
	account := jsonAuth.authByLogin(ctx, r)
	if account != nil {
		return account, nil
	}
	return jsonAuth.authByIp(ctx, r), nil
}

func (jsonAuth *HttpAuthenticator) authByLogin(ctx context.Context, r *http.Request) *auth.Account {
	logger := logging.LogFromContext(ctx)
	logger.Debug("Try auth by login/password")
	lp := request.GetLoginPass(r)
	if lp == nil {
		logger.Debug("Login/password is empty")
		return nil
	}
	account, ok := jsonAuth.accountsByLogin[lp.Login]
	logger = logger.WithField("login", lp.Login)
	if !ok {
		logger.Debug("Account not found by login")
		return nil
	}
	if account.Password != nil && *account.Password != lp.Password {
		logger.Debug("Password incorrect")
		return nil
	}
	logger.WithField("profile", account.ProfileSlug).Debug("Auth success")
	return &auth.Account{
		Login:       account.Login,
		ProfileSlug: account.ProfileSlug,
	}
}

func (jsonAuth *HttpAuthenticator) authByIp(ctx context.Context, r *http.Request) *auth.Account {
	logger := logging.LogFromContext(ctx)
	logger.Debug("Try auth by IP")
	ip := request.GetClientIpAddress(r)
	if ip == "" {
		logger.Debug("IP is empty")
		return nil
	}
	logger = logger.WithField("ip", ip)
	account, ok := jsonAuth.accountsByIp[ip]
	if !ok {
		logger.Debug("Account not found by IP")
		return nil
	}
	logger.WithField("profile", account.ProfileSlug).Debug("Auth success")
	return &auth.Account{
		Login:       account.Login,
		ProfileSlug: account.ProfileSlug,
	}
}
