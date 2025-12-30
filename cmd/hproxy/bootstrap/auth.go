package bootstrap

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/loader"
)

func InitAuth(config *Config) (auth.HttpAuthenticator, error) {
	log.Info("Creating authenticator")
	if config.JsonAuth != "" {
		auth, err := loader.LoadAuthRepository(config.JsonAuth)
		if err != nil {
			return nil, fmt.Errorf("failed to bootstrap authenticator from json: %w", err)
		}
		return auth, nil
	}
	return nil, fmt.Errorf("No auth config")
}
