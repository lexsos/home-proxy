package request

import (
	"net"
	"net/http"
)

func GetDstDomain(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		return r.Host
	}
	return host
}
