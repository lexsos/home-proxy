package inmemory

import (
	"fmt"

	"github.com/lexsos/home-proxy/profiles"
)

func (repo *InMemoryProfilesRepository) GetProfile(slug string) (*profiles.ProfileConfig, error) {
	profile, ok := repo.profiles[slug]
	if !ok {
		return nil, fmt.Errorf("Profile '%s' not found", slug)
	}
	currentTime := profiles.NowTimeInLocation(profile.tz)
	for _, timeRange := range profile.timeRanges {
		if (currentTime >= timeRange.startAt) && (currentTime <= timeRange.endAt) {
			return &profiles.ProfileConfig{
				Policy:      timeRange.policy,
				DomainsSets: timeRange.domainsSets,
			}, nil
		}
	}
	return nil, fmt.Errorf("Time range not found in profile '%s'", slug)
}
