package inmemory

import (
	"time"

	"github.com/lexsos/home-proxy/internal/profiles"
	"github.com/lexsos/home-proxy/internal/profiles/times"
)

type TimeRange struct {
	policy      profiles.ProfilePolicy
	domainsSets []string
	startAt     times.DayTime
	endAt       times.DayTime
	weekDays    map[time.Weekday]struct{}
}

type Profile struct {
	slug       string
	tz         *time.Location
	timeRanges []TimeRange
}

type InMemoryProfilesRepository struct {
	profiles map[string]Profile
}
