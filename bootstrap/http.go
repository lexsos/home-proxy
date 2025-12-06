package bootstrap

import (
	"net/http"
	"time"

	"github.com/lexsos/home-proxy/handlers"
)

func InitHttpServer(config *Config, httPproxyHandler *handlers.HttpProxyHandler) (*http.Server, error) {
	server := &http.Server{
		Addr:         config.ProxyAddr,
		Handler:      http.HandlerFunc(httPproxyHandler.Handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server, nil
}
