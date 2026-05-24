package bootstrap

import (
	"errors"
	"flag"
	"fmt"
	"maps"
	"slices"
)

type LogFormat string

const (
	LogFormatText LogFormat = "text"
	LogFormatJson LogFormat = "json"
)

var availableLogFormats = map[LogFormat]bool{
	LogFormatText: true,
	LogFormatJson: true,
}

type Config struct {
	ProxyAddr string
	SocksAddr string
	AuthSocksAddr string
	LogLevel  string
	LogFormat LogFormat
	JsonAuth  string
}

func ParseConfig() (*Config, error) {
	proxyAddr := flag.String("proxy-addr", "", "Http proxy address")
	socksAddr := flag.String("socks-addr", "", "Socks proxy address")
	authSocksAddr := flag.String("auth-socks-addr", "", "Socks proxy with auth address")
	logLevel := flag.String("log-level", "info", "Log level")
	logFormat := flag.String("log-format", string(LogFormatText), "Log format")
	jsonAuth := flag.String("auth-file", "", "Json file with auth data")
	flag.Parse()

	if *proxyAddr == "" && *socksAddr == "" && *authSocksAddr == "" {
		return nil, errors.New("http or socks address is required")
	}
	if !availableLogFormats[LogFormat(*logFormat)] {
		formats := slices.Collect(maps.Keys(availableLogFormats))
		return nil, fmt.Errorf("unknown log format: '%s', should be one of: %v", *logFormat, formats)
	}
	config := Config{
		ProxyAddr: *proxyAddr,
		SocksAddr: *socksAddr,
		AuthSocksAddr: *authSocksAddr,
		LogLevel:  *logLevel,
		LogFormat: LogFormat(*logFormat),
		JsonAuth:  *jsonAuth,
	}
	return &config, nil
}
