package response

import (
	"net/http"
)

func RequireAuth(w http.ResponseWriter) {
	w.Header().Set("Proxy-Authenticate", `Basic realm="Proxy"`)
	http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
}
