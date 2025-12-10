package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/domains"
	domainsInmemory "github.com/lexsos/home-proxy/domains/inmemory"
	profilesInmemory "github.com/lexsos/home-proxy/profiles/inmemory"
)

func InitDomainMatcher(config *Config) (domains.DomainMatcher, error) {
	if config.JsonAuth != "" {
		return domainsInmemory.NewDomainSetRepositoryFromJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}

func InitProfileRepository(config *Config) (*profilesInmemory.InMemoryProfilesRepository, error) {
	if config.JsonAuth != "" {
		return profilesInmemory.NewProfilesRepositoryFronJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}
