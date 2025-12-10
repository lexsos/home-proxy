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
	filter, err := bootstrap.InitFilter(config)
	log.Print(err)
	log.Print(filter.HasAccess("admin", "api.yandex.ru"))
	log.Print(filter.HasAccess("admin", "youtube.com"))
	log.Print(filter.HasAccess("children", "api.yandex.ru"))
	log.Print(filter.HasAccess("children", "youtube.com"))
	log.WithFields(log.Fields{"addres": config.ProxyAddr}).Info("Starting HTTPS proxy")
	log.Fatal(server.ListenAndServe())
}
