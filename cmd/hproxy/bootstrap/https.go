package bootstrap

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/https"
)

func InitHttpServer(config *Config, authenticator auth.HttpAuthenticator, reqFilter *filters.RequestFilter) (*http.Server, error) {
	log.Info("Creating http server")
	httPproxyHandler := https.NewHttpProxyHandler(authenticator, reqFilter)
	server := &http.Server{
		Addr:         config.ProxyAddr,
		Handler:      http.HandlerFunc(httPproxyHandler.Handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server, nil
}
