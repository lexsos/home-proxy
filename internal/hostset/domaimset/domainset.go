package domaimset

type DomainSet interface {
	Contains(domain string) (bool, error)
}

type MatchType int

const (
	ExactDomain MatchType = 1
	SubDomains  MatchType = 2
)
