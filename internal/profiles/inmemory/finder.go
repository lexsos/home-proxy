package inmemory

import (
	"context"
	"fmt"

	"github.com/lexsos/home-proxy/internal/profiles"
	"github.com/lexsos/home-proxy/pkg/logging"
)

func (repo *InMemoryProfilesRepository) GetProfile(ctx context.Context, slug string) (*profiles.ProfileConfig, error) {
	logger := logging.LogFromContext(ctx)
	profile, ok := repo.profiles[slug]
	if !ok {
		return nil, fmt.Errorf("Profile '%s' not found", slug)
	}
	current := profiles.NowTimeInLocation(profile.tz)
	for i, timeRange := range profile.timeRanges {
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
		logger.Debugf("Use time range %d from profile '%s'", i, profile.slug)
		return &profiles.ProfileConfig{
			Policy:      timeRange.policy,
			DomainsSets: timeRange.domainsSets,
		}, nil

	}
	return nil, fmt.Errorf("Time range not found in profile '%s'", slug)
}
