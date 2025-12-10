package bootstrap

import (
	"fmt"

	"github.com/lexsos/home-proxy/domains"
	domainsInmemory "github.com/lexsos/home-proxy/domains/inmemory"
	"github.com/lexsos/home-proxy/filters"
	"github.com/lexsos/home-proxy/profiles"
	profilesInmemory "github.com/lexsos/home-proxy/profiles/inmemory"
)

func InitDomainMatcher(config *Config) (domains.DomainMatcher, error) {
	if config.JsonAuth != "" {
		return domainsInmemory.NewDomainSetRepositoryFromJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}

func InitProfileRepository(config *Config) (profiles.ProfilesRepository, error) {
	if config.JsonAuth != "" {
		return profilesInmemory.NewProfilesRepositoryFronJson(config.JsonAuth)
	}
	return nil, fmt.Errorf("No filters config")
}

func InitFilter(config *Config) (*filters.RequestFilter, error) {
	domains, err := InitDomainMatcher(config)
	if err != nil {
		return nil, err
	}
	profileRepo, err := InitProfileRepository(config)
	if err != nil {
		return nil, err
	}
	return filters.NewRequestFilter(profileRepo, domains), nil
}
