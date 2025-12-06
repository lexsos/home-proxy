package main

import (
	"fmt"
	"github.com/lexsos/home-proxy/bootstrap"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := bootstrap.ParseConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hello, world!")
	fmt.Println(config)
}
