package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/auth/json"
)

func InitAuth(config *Config) (auth.HttpAuthenticator, error) {
	if config.JsonAuth != "" {
		return json.NewJsonHttpAuthenticator(config.JsonAuth)
	}
	return nil, fmt.Errorf("No auth config")
}
