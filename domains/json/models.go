package json

type MatchType string

const (
	Exact      MatchType = "exact"
	SubDomains MatchType = "subdomains"
)

type Domain struct {
	Dns  string
	Type MatchType
}

type DomainsSet struct {
	Slug    string
	Domains map[string]Domain
}

type DomainSetRepository struct {
	domainsSets map[string]DomainsSet
}
