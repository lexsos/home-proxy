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
	current := profiles.NowTimeInLocation(profile.tz)
	for _, timeRange := range profile.timeRanges {
		// Too early
		if current.Time < timeRange.startAt {
			continue
		}
		// Too late
		if timeRange.endAt < current.Time {
			continue
		}
		// Improper day of week
		if _, ok := timeRange.weekDays[current.Day]; !ok && len(timeRange.weekDays) > 0 {
			continue
		}
		return &profiles.ProfileConfig{
			Policy:      timeRange.policy,
			DomainsSets: timeRange.domainsSets,
		}, nil

	}
	return nil, fmt.Errorf("Time range not found in profile '%s'", slug)
}
