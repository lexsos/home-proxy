package profiles

type ProfilesFilter interface {
	IsAllow(slug string, domain string) (bool, error)
}
