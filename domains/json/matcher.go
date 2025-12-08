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
	data := prepareDomain(domain)
	for _, setSlug := range sets {
		// Get domain set by slug, if not found continue loop
		domainSet, ok := repo.domainsSets[setSlug]
		if !ok {
			continue
		}
		// Try check original DNS for domain in current domains set
		domain, ok := domainSet.Domains[data.original]
		if ok && (domain.Type == Exact || domain.Type == SubDomains) {
			return true, nil
		}
		// Try check parent domains in current domains set
		for _, parent := range data.parents {
			domain, ok := domainSet.Domains[parent]
			if !ok {
				continue
			}
			if domain.Type == SubDomains {
				return true, nil
			}
		}
	}
	return false, nil
}
