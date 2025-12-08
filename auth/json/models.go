package json

type JsonHttpAuthenticator struct {
	accountsByLogin map[string]JsonAccount
	accountsByIp    map[string]JsonAccount
}
