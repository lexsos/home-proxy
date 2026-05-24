package main

import (
	"net/http"

	"github.com/armon/go-socks5"

	"github.com/lexsos/home-proxy/cmd/hproxy/bootstrap"
)

type Deps struct {
	config      *bootstrap.Config
	httpServer  *http.Server
	socksServer *socks5.Server
	authSocksServer *socks5.Server
}

func NewDeps() (*Deps, error) {
	config, err := bootstrap.ParseConfig()
	if err != nil {
		return nil, err
	}
	bootstrap.InitLog(config)
	authenticator, err := bootstrap.InitAuth(config)
	if err != nil {
		return nil, err
	}
	filter, err := bootstrap.InitFilter(config)
	if err != nil {
		return nil, err
	}

	deps := &Deps{
		config:      config,
		httpServer:  nil,
		socksServer: nil,
	}
	if config.ProxyAddr != "" {
		httpServer, err := bootstrap.InitHttpServer(config, authenticator, filter)
		if err != nil {
			return nil, err
		}
		deps.httpServer = httpServer
	}
	if config.SocksAddr != "" {
		socksServer, err := bootstrap.InitSocksServer(filter, authenticator, true)
		if err != nil {
			return nil, err
		}
		deps.socksServer = socksServer
	}
	if config.AuthSocksAddr != "" {
		authSocksServer, err := bootstrap.InitSocksServer(filter, authenticator, false)
		if err != nil {
			return nil, err
		}
		deps.authSocksServer = authSocksServer
	}
	return deps, nil
}

func (d *Deps) HasHttpServer() bool {
	return d.httpServer != nil
}

func (d *Deps) HasSocksServer() bool {
	return d.socksServer != nil
}

func (d *Deps) HasAuthSocksServer() bool {
	return d.authSocksServer != nil
}
