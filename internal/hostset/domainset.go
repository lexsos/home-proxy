package hostset

type DomainSet interface {
	Contains(domain string) (bool, error)
}
