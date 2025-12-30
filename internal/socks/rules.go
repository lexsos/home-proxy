package socks

import (
	"context"

	"github.com/armon/go-socks5"

	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/utils/logging"
)

var allowedMethods = map[uint8]struct{}{
	socks5.NoAuth:       {},
	socks5.UserPassAuth: {},
}

type SocksRules struct {
	filter filters.RequestFilter
}

func NewSocksRules(filter filters.RequestFilter) *SocksRules {
	return &SocksRules{
		filter: filter,
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
	_, hasAith := allowedMethods[req.AuthContext.Method]
	if !hasAith {
		logger.Debugf("invalid auth method %d for '%s'", req.AuthContext.Method, req.RemoteAddr.IP)
		return ""
	}
	profile, ok := req.AuthContext.Payload[profileField]
	if !ok {
		logger.Debugf("profile not found for '%s'", req.RemoteAddr.IP)
		return ""
	}
	return profile
}

func (rules *SocksRules) dst(req *socks5.Request) string {
	dest := req.DestAddr.FQDN
	if dest != "" {
		return dest
	}
	return req.DestAddr.IP.String()
}
