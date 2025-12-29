package domainset

import "strings"

type DomainSig struct {
	original string
	parents  []string
}

func NewDomainSig(domain string) *DomainSig {
	original := strings.ToLower(domain)

	var parents []string
	parts := strings.Split(original, ".")
	for i := 1; i < len(parts); i++ {
		parent := strings.Join(parts[i:], ".")
		parents = append(parents, parent)
	}

	return &DomainSig{
		original: original,
		parents:  parents,
	}
}

func (d *DomainSig) Original() string {
	return d.original
}

func (d *DomainSig) Parents() []string {
	return d.parents
}
