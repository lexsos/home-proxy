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
	server, err := bootstrap.NewHttpProxy(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("Starting HTTPS proxy on", config.ProxyAddr)
	log.Fatal(server.ListenAndServe())
}
