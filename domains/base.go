package domains

type DomainMatcher interface {
	Match(domain string, setsSlugs []string) bool
}
