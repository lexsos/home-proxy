package domaimset

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
