package bootstrap

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/handlers"
)

func InitHttpServer(config *Config, authenticator auth.HttpAuthenticator) (*http.Server, error) {
	log.Info("Creating http server")
	httPproxyHandler := handlers.NewProxyHandler(authenticator)
	server := &http.Server{
		Addr:         config.ProxyAddr,
		Handler:      http.HandlerFunc(httPproxyHandler.Handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server, nil
}
