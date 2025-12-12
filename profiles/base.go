package profiles

import "context"

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
	GetProfile(ctx context.Context, slug string) (*ProfileConfig, error)
}
