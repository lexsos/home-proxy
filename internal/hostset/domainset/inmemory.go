package domainset

import (
	"strings"
)

type InMemoryDomainsSet struct {
	domains map[string]MatchType
}

func NewInMemoryDomainsSet() *InMemoryDomainsSet {
	return &InMemoryDomainsSet{
		domains: make(map[string]MatchType),
	}
}

func (d *InMemoryDomainsSet) Add(domain string, mtype MatchType) {
	domain = strings.ToLower(domain)
	d.domains[domain] = mtype
}

func (d *InMemoryDomainsSet) Contains(domain string) (bool, error) {
	return d.ContainsSig(NewDomainSig(domain))
}

func (d *InMemoryDomainsSet) ContainsSig(sig *DomainSig) (bool, error) {
	mtype, ok := d.domains[sig.Original()]
	if ok && (mtype == ExactDomain || mtype == SubDomains) {
		return true, nil
	}
	for _, parent := range sig.Parents() {
		mtype, ok := d.domains[parent]
		if !ok || mtype != SubDomains {
			continue
		}
		return true, nil
	}
	return false, nil
}
