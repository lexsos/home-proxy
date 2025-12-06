package bootstrap

import (
	"errors"
	"flag"
)

type Config struct {
	ProxyAddr string
	LogLevel  string
	JsonAuth  string
}

func ParseConfig() (*Config, error) {
	proxyAddr := flag.String("proxy-addr", "", "Proxy address")
	logLevel := flag.String("log-level", "info", "Log level")
	jsonAuth := flag.String("auth-file", "", "Json file with auth data")
	flag.Parse()

	if *proxyAddr == "" {
		return nil, errors.New("empty proxy address")
	}
	config := Config{
		ProxyAddr: *proxyAddr,
		LogLevel:  *logLevel,
		JsonAuth:  *jsonAuth,
	}
	return &config, nil
}
