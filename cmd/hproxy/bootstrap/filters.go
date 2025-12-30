package bootstrap

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/hostset"
	"github.com/lexsos/home-proxy/internal/loader"
	"github.com/lexsos/home-proxy/internal/profiles"
	profilesInmemory "github.com/lexsos/home-proxy/internal/profiles/inmemory"
)

func InitHostRepository(config *Config) (hostset.HostRepository, error) {
	log.Info("Loading hosts")
	if config.JsonAuth != "" {
		repo, err := loader.LoadHostRepository(config.JsonAuth)
		if err != nil {
			return nil, fmt.Errorf("failed to bootstrap hosts from json: %w", err)
		}
		return repo, nil
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
	hosts, err := InitHostRepository(config)
	if err != nil {
		return nil, err
	}
	profileRepo, err := InitProfileRepository(config)
	if err != nil {
		return nil, err
	}
	return filters.NewRequestFilter(profileRepo, hosts), nil
}
