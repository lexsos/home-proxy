package filters

import (
	"fmt"

	"github.com/lexsos/home-proxy/domains"
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

func (filter *RequestFilter) HasAccess(profileSlug string, domain string) (bool, error) {
	cfg, err := filter.profilesRepo.GetProfile(profileSlug)
	if err != nil {
		return false, fmt.Errorf("fail extract profile cfg: %w", err)
	}
	switch cfg.Policy {
	case profiles.Allow:
		return true, nil
	case profiles.Strict:
		isMatch, err := filter.domainMatcher.Match(domain, cfg.DomainsSets)
		if err != nil {
			return false, fmt.Errorf("fail match domain '%s': %w", domain, err)
		}
		return isMatch, nil
	default:
		return false, fmt.Errorf("Unknown policy '%s' for profile '%s'", cfg.Policy, profileSlug)
	}
}
