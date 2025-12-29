package filters

import (
	"context"
	"fmt"

	"github.com/lexsos/home-proxy/internal/hostset"
	"github.com/lexsos/home-proxy/internal/profiles"
	"github.com/lexsos/home-proxy/internal/utils/logging"
)

type RequestFilter struct {
	profilesRepo profiles.ProfilesRepository
	hostsRepo    hostset.HostRepository
}

func NewRequestFilter(profilesRepo profiles.ProfilesRepository, hostsRepo hostset.HostRepository) *RequestFilter {
	return &RequestFilter{
		profilesRepo: profilesRepo,
		hostsRepo:    hostsRepo,
	}
}

func (filter *RequestFilter) HasAccess(ctx context.Context, profileSlug string, domain string) (bool, error) {
	logger := logging.LogFromContext(ctx)
	cfg, err := filter.profilesRepo.GetProfile(ctx, profileSlug)
	if err != nil {
		return false, fmt.Errorf("fail extract profile cfg: %w", err)
	}
	switch cfg.Policy {
	case profiles.Allow:
		logger.Debug("Use allow policy")
		return true, nil
	case profiles.Strict:
		logger.Debugf("Use strict policy with domains set: %v ", cfg.DomainsSets)
		contains, err := filter.hostsRepo.Contains(domain, cfg.DomainsSets)
		if err != nil {
			return false, fmt.Errorf("fail match domain '%s': %w", domain, err)
		}
		return contains, nil
	default:
		return false, fmt.Errorf("Unknown policy '%s' for profile '%s'", cfg.Policy, profileSlug)
	}
}
