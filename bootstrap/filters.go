package bootstrap

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/domains"
	domainsInmemory "github.com/lexsos/home-proxy/domains/inmemory"
	"github.com/lexsos/home-proxy/filters"
	"github.com/lexsos/home-proxy/profiles"
	profilesInmemory "github.com/lexsos/home-proxy/profiles/inmemory"
)

func InitDomainMatcher(config *Config) (domains.DomainMatcher, error) {
	log.Info("Loading domains")
	if config.JsonAuth != "" {
		domains, err := domainsInmemory.NewDomainSetRepositoryFromJson(config.JsonAuth)
		if err != nil {
			return nil, fmt.Errorf("failed to bootstrap domains from json: %w", err)
		}
		return domains, nil
	}
	return nil, fmt.Errorf("No filters config")
}

func InitProfileRepository(config *Config) (profiles.ProfilesRepository, error) {
	log.Info("Loading profiles")
	if config.JsonAuth != "" {
		profiles, err := profilesInmemory.NewProfilesRepositoryFronJson(config.JsonAuth)
		if err != nil {
			return nil, fmt.Errorf("failed to bootstrap profiles from json: %w", err)
		}
		return profiles, nil
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
