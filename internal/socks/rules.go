package socks

import (
	"context"

	"github.com/armon/go-socks5"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/utils/logging"
)

type SocksRules struct {
	filter        *filters.RequestFilter
	authenticator auth.Authenticator
}

func NewSocksRules(filter *filters.RequestFilter, authenticator auth.Authenticator) *SocksRules {
	return &SocksRules{
		filter:        filter,
		authenticator: authenticator,
	}
}

func (rules *SocksRules) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	logger := logging.LogFromContext(ctx)
	profile := rules.profile(ctx, req)
	if profile == "" {
		logger.Debugf("unknown profile for '%s'", req.RemoteAddr.IP)
		return ctx, false
	}
	dst := rules.dst(req)
	hasAccess, err := rules.filter.HasAccess(ctx, profile, dst)
	if err != nil {
		logger.Errorf("fail check access in socks rules: %s", err)
		return ctx, false
	}
	if hasAccess {
		logger.Debugf("allow access for '%s' to '%s' in socks rules", profile, dst)
		return ctx, true
	}
	logger.Debugf("deny access for '%s' to '%s' in socks rules", profile, dst)
	return ctx, hasAccess
}

func (rules *SocksRules) profile(ctx context.Context, req *socks5.Request) string {
	logger := logging.LogFromContext(ctx)
	switch req.AuthContext.Method {
	case socks5.NoAuth:
		src := req.RemoteAddr.IP.String()
		logger.Debugf("auth by ip '%s' in socks rules", src)
		account, err := rules.authenticator.AuthByIp(ctx, src)
		if err != nil {
			logger.Errorf("fail auth by ip in socks rules: %s", err)
			return ""
		}
		if account == nil {
			logger.Debugf("unknown account for '%s'", src)
			return ""
		}
		logger.Debugf("success auth by ip for '%s'", account.Login)
		return account.ProfileSlug
	case socks5.UserPassAuth:
		logger.Debugf("auth by login in socks rules")
		login, ok := req.AuthContext.Payload["Username"]
		if !ok {
			logger.Debugf("login not found for '%s'", req.RemoteAddr.IP)
			return ""
		}
		account, err := rules.authenticator.GetByLogin(login)
		if err != nil {
			logger.Errorf("fail get account by login in socks rules: %s", err)
			return ""
		}
		if account == nil {
			logger.Debugf("unknown account for '%s'", login)
			return ""
		}
		logger.Debugf("success auth by login for '%s'", account.Login)
		return account.ProfileSlug
	default:
		logger.Debugf("unknown auth method %d for '%s'", req.AuthContext.Method, req.RemoteAddr.IP)
		return ""
	}
}

func (rules *SocksRules) dst(req *socks5.Request) string {
	dest := req.DestAddr.FQDN
	if dest != "" {
		return dest
	}
	return req.DestAddr.IP.String()
}
