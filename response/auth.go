package response

import (
	"fmt"
	"net/http"
)

func RequireAuth(w http.ResponseWriter) {
	w.Header().Set("Proxy-Authenticate", `Basic realm="Proxy"`)
	http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
}

func DomainForbidden(w http.ResponseWriter, domain string) {
	massage := fmt.Sprintf("Domain '%s' forbidden", domain)
	http.Error(w, massage, http.StatusForbidden)
}
