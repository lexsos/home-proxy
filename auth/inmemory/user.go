package inmemory

import (
	"net/http"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/request"
)

func (jsonAuth *HttpAuthenticator) GetUser(r *http.Request) (*auth.Account, error) {
	account := jsonAuth.authByLogin(r)
	if account != nil {
		return account, nil
	}
	return jsonAuth.authByIp(r), nil
}

func (jsonAuth *HttpAuthenticator) authByLogin(r *http.Request) *auth.Account {
	lp := request.GetLoginPass(r)
	if lp == nil {
		return nil
	}
	account, ok := jsonAuth.accountsByLogin[lp.Login]
	if !ok {
		return nil
	}
	if account.Password != nil && *account.Password != lp.Password {
		return nil
	}
	return &auth.Account{
		Login:       account.Login,
		ProfileSlug: account.ProfileSlug,
	}
}

func (jsonAuth *HttpAuthenticator) authByIp(r *http.Request) *auth.Account {
	ip := request.GetIpAddress(r)
	if ip == "" {
		return nil
	}
	account, ok := jsonAuth.accountsByIp[ip]
	if !ok {
		return nil
	}
	return &auth.Account{
		Login:       account.Login,
		ProfileSlug: account.ProfileSlug,
	}
}
