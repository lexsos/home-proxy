package inmemory

type AccountData struct {
	Login       string   `json:"login"`
	Password    *string  `json:"password"`
	Ips         []string `json:"ips"`
	ProfileSlug string   `json:"profile"`
}

type HttpAuthenticator struct {
	accountsByLogin map[string]AccountData
	accountsByIp    map[string]AccountData
}
