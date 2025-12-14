package inmemory

import (
	"time"

	"github.com/lexsos/home-proxy/internal/profiles"
)

type TimeRange struct {
	policy      profiles.ProfilePolicy
	domainsSets []string
	startAt     profiles.DayTime
	endAt       profiles.DayTime
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
