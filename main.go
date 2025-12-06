package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/bootstrap"
	"github.com/lexsos/home-proxy/handlers"
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
	httPproxyHandler := handlers.NewProxyHandler(authenticator)
	server, err := bootstrap.InitHttpServer(config, httPproxyHandler)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.WithFields(log.Fields{"addres": config.ProxyAddr}).Info("Starting HTTPS proxy")
	log.Fatal(server.ListenAndServe())
}
