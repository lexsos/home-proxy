package bootstrap

import (
	"net/http"
	"time"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/handlers"
)

func InitHttpServer(config *Config, authenticator auth.HttpAuthenticator) (*http.Server, error) {
	httPproxyHandler := handlers.NewProxyHandler(authenticator)
	server := &http.Server{
		Addr:         config.ProxyAddr,
		Handler:      http.HandlerFunc(httPproxyHandler.Handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server, nil
}
