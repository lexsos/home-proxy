package profiles

type ProfilePolicy string

const (
	Allow  ProfilePolicy = "allow"
	Strict ProfilePolicy = "strict"
)

type ProfileConfig struct {
	Policy      ProfilePolicy
	DomainsSets []string
}

type ProfilesRepository interface {
	GetProfile(slug string) (*ProfileConfig, error)
}
