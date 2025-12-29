package domaimset

type MatchType int

const (
	ExactDomain MatchType = 1
	SubDomains  MatchType = 2
)

type DomainSet interface {
	Contains(domain string) (bool, error)
	ContainsSig(sig *DomainSig) (bool, error)
}
