package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	deps, err := NewDeps()
	if err != nil {
		log.Fatalf("fail create deps: %s", err)
		return
	}

	errChan := make(chan error, 2)
	if deps.HasHttpServer() {
		go runHttp(deps, errChan)
	}
	if deps.HasSocksServer() {
		go runSocks(deps, errChan)
	}
	if deps.HasAuthSocksServer() {
		go runAuthSocks(deps, errChan)
	}
	log.Fatal(<-errChan)
}

func runHttp(deps *Deps, errChan chan<- error) {
	log.Infof("Starting HTTP/HTTPS proxy on port %s", deps.config.ProxyAddr)
	errChan <- deps.httpServer.ListenAndServe()
}

func runSocks(deps *Deps, errChan chan<- error) {
	log.Infof("Starting SOCKS5 proxy on port %s", deps.config.SocksAddr)
	errChan <- deps.socksServer.ListenAndServe("tcp", deps.config.SocksAddr)
}

func runAuthSocks(deps *Deps, errChan chan<- error) {
	log.Infof("Starting SOCKS5 with auth proxy on port %s", deps.config.AuthSocksAddr)
	errChan <- deps.authSocksServer.ListenAndServe("tcp", deps.config.AuthSocksAddr)
}
