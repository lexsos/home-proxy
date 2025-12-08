package json

import "strings"

type parsedDomain struct {
	original string
	parents  []string
}

func prepareDomain(domain string) *parsedDomain {
	original := strings.ToLower(domain)

	var parents []string
	parts := strings.Split(original, ".")
	for i := 1; i < len(parts); i++ {
		parent := strings.Join(parts[i:], ".")
		parents = append(parents, parent)
	}

	return &parsedDomain{
		original: original,
		parents:  parents,
	}
}

func (repo *DomainSetRepository) Match(domain string, sets []string) (bool, error) {
	return false, nil
}
