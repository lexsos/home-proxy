package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/domains"
	"github.com/lexsos/home-proxy/domains/inmemory"
)

func InitDomainMatcher(config *Config) (domains.DomainMatcher, error) {
	if config.JsonAuth != "" {
		return inmemory.NewDomainSetRepositoryFromJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}
