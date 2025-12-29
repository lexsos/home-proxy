package hostset

import (
	"fmt"
	"net"

	"github.com/lexsos/home-proxy/internal/hostset/domainset"
	"github.com/lexsos/home-proxy/internal/hostset/ipset"
)

type InMemoryHostRepository struct {
	hostSets map[string]*HostSet
}

func NewInMemoryHostRepository() *InMemoryHostRepository {
	return &InMemoryHostRepository{
		hostSets: make(map[string]*HostSet),
	}
}

func (repo *InMemoryHostRepository) AddHostSet(name string, ips ipset.IpSet, domans domainset.DomainSet) {
	repo.hostSets[name] = NewHostSet(domans, ips)
}

func (repo *InMemoryHostRepository) Contains(host string, setNames []string) (bool, error) {
	ip := net.ParseIP(host)
	if ip == nil {
		return repo.containsDomainSig(domainset.NewDomainSig(host), setNames)
	}
	ipSig, err := ipset.NewIpSignature(ip)
	if err != nil {
		return false, fmt.Errorf("invalid ip: %s", host)
	}
	return repo.containsIpSig(ipSig, setNames)
}

func (repo *InMemoryHostRepository) containsIpSig(ipSig *ipset.IpSignature, setNames []string) (bool, error) {
	for _, setName := range setNames {
		hostSet, ok := repo.hostSets[setName]
		if !ok {
			continue
		}
		contains, err := hostSet.ContainsIpSig(ipSig)
		if err != nil {
			return false, fmt.Errorf("can't check ip %s in set %s: %w", ipSig.Src(), setName, err)
		}
		if contains {
			return true, nil
		}
	}
	return false, nil
}

func (repo *InMemoryHostRepository) containsDomainSig(domainSig *domainset.DomainSig, setNames []string) (bool, error) {
	for _, setName := range setNames {
		hostSet, ok := repo.hostSets[setName]
		if !ok {
			continue
		}
		contains, err := hostSet.ContainsDomainSig(domainSig)
		if err != nil {
			return false, fmt.Errorf("can't check domain %s in set %s: %w", domainSig.Original(), setName, err)
		}
		if contains {
			return true, nil
		}
	}
	return false, nil
}
