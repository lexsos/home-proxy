package main

import (
	"net/http"

	"github.com/armon/go-socks5"
	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/cmd/hproxy/bootstrap"
)

func main() {
	config, err := bootstrap.ParseConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bootstrap.InitLog(config)
	authenticator, err := bootstrap.InitAuth(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	filter, err := bootstrap.InitFilter(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	httpServer, err := bootstrap.InitHttpServer(config, authenticator, filter)
	if err != nil {
		log.Fatal(err)
		return
	}
	socksServer, err := bootstrap.InitSocksServer(filter, authenticator)
	if err != nil {
		log.Fatal(err)
		return
	}
	runHttp(config, httpServer)
	runSocks(config, socksServer)
}

func runHttp(config *bootstrap.Config, server *http.Server) {
	log.Infof("Starting HTTP/HTTPS proxy on port %s", config.ProxyAddr)
	log.Fatal(server.ListenAndServe())
}

func runSocks(config *bootstrap.Config, server *socks5.Server) {
	log.Infof("Starting SOCKS5 proxy on port %s", config.SocksAddr)
	log.Fatal(server.ListenAndServe("tcp", config.SocksAddr))
}
