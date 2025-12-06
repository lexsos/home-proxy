package handlers

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/response"
)

type HttpProxyHandler struct {
	authenticator auth.HttpAuthenticator
}

func NewProxyHandler(authenticator auth.HttpAuthenticator) *HttpProxyHandler {
	return &HttpProxyHandler{
		authenticator: authenticator,
	}
}

func (proxy *HttpProxyHandler) Handler(w http.ResponseWriter, r *http.Request) {
	account, err := proxy.authenticator.GetUser(r)
	if err != nil {
		log.Warn("Auth fail: ", err)
		response.RequireAuth(w)
		return
	}
	if account == nil {
		response.RequireAuth(w)
		log.Info("Fail auth for: ", r.RemoteAddr)
		return
	}
	log.Info("New connection from ", r.RemoteAddr, " to ", r.Host)
	if strings.ToLower(r.Method) == "connect" {
		handleTunnel(w, r)
	} else {
		handleHTTPProxy(w, r)
	}
}
