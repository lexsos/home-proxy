package domains

type jsonDomain struct {
	Dns  string    `json:"dns"`
	Type MatchType `json:"type"`
}

type jsonDomainSet struct {
	Slug    string       `json:"slug"`
	Domains []jsonDomain `json:"ips"`
}

type jsonConfig struct {
	DomainsSets []jsonDomainSet `json:"domains_sets"`
}

func NewDomainSetRepositoryFromJson(fileName string) (*DomainSetRepository, error) {
	return &DomainSetRepository{}, nil
}
