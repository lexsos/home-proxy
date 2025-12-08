package json

type JsonAccount struct {
	Login       string   `json:"login"`
	Password    *string  `json:"password"`
	Ips         []string `json:"ips"`
	ProfileSlug string   `json:"profile"`
}

type JsonAccounts struct {
	Accounts []JsonAccount `json:"accounts"`
}

type JsonHttpAuthenticator struct {
	accountsByLogin map[string]JsonAccount
	accountsByIp    map[string]JsonAccount
}
