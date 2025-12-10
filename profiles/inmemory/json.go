package inmemory

import "github.com/lexsos/home-proxy/profiles"

type jsonTimeRange struct {
	policy      profiles.ProfilePolicy `json:"accounts"`
	domainsSets []string               `json:"domains_sets"`
	startAt     string                 `json:"start_at"`
	endAt       string                 `json:"end_at"`
}

type jsonProfile struct {
	slug   string          `json:"slug"`
	tz     string          `json:"tz"`
	ranges []jsonTimeRange `json:"ranges"`
}

type jsonProfilesList struct {
	profiles []jsonProfile `json:"profiles"`
}

func NewProfilesRepositoryFronJson(fileName string) (*InMemoryProfilesRepository, error) {
	return nil, nil
}
