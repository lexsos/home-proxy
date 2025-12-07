package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/bootstrap"
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
	server, err := bootstrap.InitHttpServer(config, authenticator)
	if err != nil {
		log.Fatal(err)
		return
	}
	domains, err := bootstrap.InitDomainMatcher(config)
	log.Print(domains, err)
	log.WithFields(log.Fields{"addres": config.ProxyAddr}).Info("Starting HTTPS proxy")
	log.Fatal(server.ListenAndServe())
}
