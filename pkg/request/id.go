package request

import (
	"net/http"

	"github.com/google/uuid"
)

func GetOrGenId(r *http.Request) string {
	id := r.Header.Get("Request-Id")
	if id != "" {
		return id
	}
	return uuid.New().String()
}
