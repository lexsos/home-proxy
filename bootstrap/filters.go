package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/domains/json"
)

func InitDomainMatcher(config *Config) (*json.DomainSetRepository, error) {
	if config.JsonAuth != "" {
		return json.NewDomainSetRepository(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}
