package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type LiginPass struct {
	Login    string
	Password string
}

func getLoginPass(r *http.Request) *LiginPass {
	auth := r.Header.Get("Proxy-Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Basic ") {
		return nil
	}
	payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return nil
	}
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil
	}
	return &LiginPass{
		Login:    pair[0],
		Password: pair[1],
	}
}
