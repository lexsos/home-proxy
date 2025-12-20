package hostset

type DomainSet interface {
	Contains(domain string) (bool, error)
}

type MatchType string

const (
	ExactDomain MatchType = "exactdomain"
	SubDomains  MatchType = "subdomains"
)

type Domain struct {
	Dns  string
	Type MatchType
}

type InMemoryDomainsSet struct {
	Domains map[string]Domain
}

func NewInMemoryDomainsSet() *InMemoryDomainsSet {
	return &InMemoryDomainsSet{
		Domains: make(map[string]Domain),
	}
}
