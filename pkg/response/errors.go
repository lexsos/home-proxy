package response

import (
	"net/http"
)

func InternalError(w http.ResponseWriter) {
	http.Error(w, "Proxy error", http.StatusInternalServerError)
}
