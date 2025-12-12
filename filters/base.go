package filters

import (
	"context"
	"fmt"

	"github.com/lexsos/home-proxy/domains"
	"github.com/lexsos/home-proxy/logging"
	"github.com/lexsos/home-proxy/profiles"
)

type RequestFilter struct {
	profilesRepo  profiles.ProfilesRepository
	domainMatcher domains.DomainMatcher
}

func NewRequestFilter(profilesRepo profiles.ProfilesRepository, domainMatcher domains.DomainMatcher) *RequestFilter {
	return &RequestFilter{
		profilesRepo:  profilesRepo,
		domainMatcher: domainMatcher,
	}
}

func (filter *RequestFilter) HasAccess(ctx context.Context, profileSlug string, domain string) (bool, error) {
	logger := logging.LogFromContext(ctx)
	cfg, err := filter.profilesRepo.GetProfile(profileSlug)
	if err != nil {
		return false, fmt.Errorf("fail extract profile cfg: %w", err)
	}
	switch cfg.Policy {
	case profiles.Allow:
		logger.Debug("Use allow policy")
		return true, nil
	case profiles.Strict:
		logger.Debug("Use strict policy with domains set: ", cfg.DomainsSets)
		isMatch, err := filter.domainMatcher.Match(domain, cfg.DomainsSets)
		if err != nil {
			return false, fmt.Errorf("fail match domain '%s': %w", domain, err)
		}
		return isMatch, nil
	default:
		return false, fmt.Errorf("Unknown policy '%s' for profile '%s'", cfg.Policy, profileSlug)
	}
}
