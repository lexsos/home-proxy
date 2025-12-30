package fields

import (
	"encoding/json"
	"fmt"
)

type MatchType string

const (
	ExactHost       MatchType = "exact"
	SubDomainsHosts MatchType = "subdomains"
	Ip              MatchType = "ip"
)

var availableMatchTypes = map[MatchType]struct{}{
	ExactHost:       {},
	SubDomainsHosts: {},
	Ip:              {},
}

func (u *MatchType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if _, ok := availableMatchTypes[MatchType(s)]; !ok {
		return fmt.Errorf("unknown match type: '%s'", s)
	}
	*u = MatchType(s)
	return nil
}
