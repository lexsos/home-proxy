package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/auth/inmemory"
)

func InitAuth(config *Config) (auth.HttpAuthenticator, error) {
	if config.JsonAuth != "" {
		return inmemory.NewHttpAuthenticatorFromJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No auth config")
}
