package handlers

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func HandleProxy(w http.ResponseWriter, r *http.Request) {
	log.Info("New connection from", r.RemoteAddr, "to", r.Method, r.Host)
	log.Printf("%s %s", r.Method, r.Host)
	if strings.ToLower(r.Method) == "connect" {
		handleTunnel(w, r)
	} else {
		handleHTTPProxy(w, r)
	}
}
