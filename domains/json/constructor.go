package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonDomain struct {
	Dns  string    `json:"dns"`
	Type MatchType `json:"type"`
}

type jsonDomainSet struct {
	Slug    string       `json:"slug"`
	Domains []jsonDomain `json:"domains"`
}

type jsonConfig struct {
	DomainsSets []jsonDomainSet `json:"domains_sets"`
}

func NewDomainSetRepository(fileName string) (*DomainSetRepository, error) {
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
	domainsSets := make(map[string]DomainsSet)

	for _, domainSet := range config.DomainsSets {
		domainsMap := make(map[string]Domain)

		for _, domain := range domainSet.Domains {
			domainsMap[domain.Dns] = Domain{
				Dns:  domain.Dns,
				Type: domain.Type,
			}
		}

		domainsSets[domainSet.Slug] = DomainsSet{
			Slug:    domainSet.Slug,
			Domains: domainsMap,
		}
	}

	return &DomainSetRepository{
		domainsSets: domainsSets,
	}, nil
}
