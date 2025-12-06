package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/auth"
)

func InitAuth(config *Config) (auth.HttpAuthenticator, error) {
	if config.JsonAuth != "" {
		return auth.NewJsonHttpAuthenticator(config.JsonAuth)
	}
	return nil, fmt.Errorf("No auth config")
}
