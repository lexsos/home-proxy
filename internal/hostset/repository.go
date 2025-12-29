package hostset

type HostRepository interface {
	Contains(host string, setNames []string) (bool, error)
}
