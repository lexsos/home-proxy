package bootstrap

import (
	"net/http"
	"time"

	"github.com/lexsos/home-proxy/handlers"
)

func NewHttpProxy(config *Config) (*http.Server, error) {
	server := &http.Server{
		Addr:         config.ProxyAddr,
		Handler:      http.HandlerFunc(handlers.HandleProxy),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server, nil
}
