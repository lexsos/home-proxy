package inmemory

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lexsos/home-proxy/internal/profiles"
)

type jsonTimeRange struct {
	Policy      profiles.ProfilePolicy `json:"policy"`
	DomainsSets []string               `json:"domains_sets"`
	StartAt     string                 `json:"start_at"`
	EndAt       string                 `json:"end_at"`
	WeekDays    []string               `json:"week_days"`
}

type jsonProfile struct {
	Slug   string          `json:"slug"`
	Tz     string          `json:"tz"`
	Ranges []jsonTimeRange `json:"ranges"`
}

type jsonConfig struct {
	Profiles []jsonProfile `json:"profiles"`
}

var WeekDaysMap = map[string]time.Weekday{
	"san": time.Sunday,
	"mon": time.Monday,
	"tue": time.Tuesday,
	"wed": time.Wednesday,
	"thu": time.Thursday,
	"fri": time.Friday,
	"sat": time.Saturday,
}

func NewProfilesRepositoryFronJson(fileName string) (*InMemoryProfilesRepository, error) {
	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse JSON
	var config jsonConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Build the repository
	profilesMap := make(map[string]Profile)

	for _, jsonProf := range config.Profiles {
		// Parse timezone
		tz, err := time.LoadLocation(jsonProf.Tz)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timezone '%s' for profile '%s': %w", jsonProf.Tz, jsonProf.Slug, err)
		}

		// Parse time ranges
		var timeRanges []TimeRange
		for i, jsonRange := range jsonProf.Ranges {
			// Parse start time
			startAt, err := profiles.ParseTime(jsonRange.StartAt)
			if err != nil {
				return nil, fmt.Errorf("failed to parse start_at '%s' for profile '%s' range %d: %w", jsonRange.StartAt, jsonProf.Slug, i, err)
			}

			// Parse end time
			endAt, err := profiles.ParseTime(jsonRange.EndAt)
			if err != nil {
				return nil, fmt.Errorf("failed to parse end_at '%s' for profile '%s' range %d: %w", jsonRange.EndAt, jsonProf.Slug, i, err)
			}

			// Parse week days
			weekDays := make(map[time.Weekday]struct{}, len(jsonRange.WeekDays))
			for _, day := range jsonRange.WeekDays {
				dayOfWeek, ok := WeekDaysMap[strings.ToLower(day)]
				if !ok {
					return nil, fmt.Errorf("fail to parse day of week '%s' for profile '%s' range %d", day, jsonProf.Slug, i)
				}
				weekDays[dayOfWeek] = struct{}{}
			}

			timeRanges = append(timeRanges, TimeRange{
				policy:      jsonRange.Policy,
				domainsSets: jsonRange.DomainsSets,
				startAt:     startAt,
				endAt:       endAt,
				weekDays:    weekDays,
			})
		}

		profilesMap[jsonProf.Slug] = Profile{
			slug:       jsonProf.Slug,
			tz:         tz,
			timeRanges: timeRanges,
		}
	}

	return &InMemoryProfilesRepository{
		profiles: profilesMap,
	}, nil
}
