package main

import (
	"fmt"

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
	fmt.Println("Hello, world!")
	fmt.Println(config)
}
