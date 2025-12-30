package bootstrap

import (
	"github.com/armon/go-socks5"
	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/socks"
)

func InitSocksServer(filter *filters.RequestFilter, authenticator auth.Authenticator) (*socks5.Server, error) {
	log.Info("Creating socks5 server")
	rules := socks.NewSocksRules(filter, authenticator)
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: authenticator,
			},
			socks5.NoAuthAuthenticator{},
		},
		Rules: rules,
	}
	return socks5.New(conf)
}
