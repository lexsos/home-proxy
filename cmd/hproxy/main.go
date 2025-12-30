package main

import (
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
	_, err = bootstrap.InitSocksServer(filter, authenticator)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("Starting HTTP/HTTPS proxy on port %s", config.ProxyAddr)
	log.Fatal(httpServer.ListenAndServe())
}
