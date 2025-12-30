package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lexsos/home-proxy/internal/hostset"
	"github.com/lexsos/home-proxy/internal/hostset/domainset"
	"github.com/lexsos/home-proxy/internal/hostset/ipset"
	"github.com/lexsos/home-proxy/internal/loader/fields"
)

type jsonHost struct {
	Host fields.NotEmptyString `json:"host"`
	Type fields.MatchType      `json:"type"`
}

type jsonHostSet struct {
	Slug  fields.NotEmptyString `json:"slug"`
	Hosts []jsonHost            `json:"hosts"`
}

type jsonHostSets struct {
	HostSets []jsonHostSet `json:"hosts_sets"`
}

func LoadHostRepository(fileName string) (hostset.HostRepository, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var hostSets jsonHostSets
	if err := json.Unmarshal(data, &hostSets); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	repo := hostset.NewInMemoryHostRepository()
	for _, hostSet := range hostSets.HostSets {
		ips := ipset.NewInMemoryIpSet()
		domains := domainset.NewInMemoryDomainsSet()
		for _, host := range hostSet.Hosts {
			switch host.Type {
			case fields.ExactHost:
				domains.Add(host.Host.String(), domainset.ExactDomain)
			case fields.SubDomainsHosts:
				domains.Add(host.Host.String(), domainset.SubDomains)
			case fields.Ip:
				err := ips.Add(host.Host.String())
				if err != nil {
					return nil, fmt.Errorf("failed to add ip to set: %w", err)
				}
			}
		}
		repo.AddHostSet(hostSet.Slug.String(), ips, domains)
	}
	return repo, nil
}
