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
	LogLevel  string
	LogFormat LogFormat
	JsonAuth  string
}

func ParseConfig() (*Config, error) {
	proxyAddr := flag.String("proxy-addr", "", "Http proxy address")
	socksAddr := flag.String("socks-addr", "", "Socks proxy address")
	logLevel := flag.String("log-level", "info", "Log level")
	logFormat := flag.String("log-format", string(LogFormatText), "Log format")
	jsonAuth := flag.String("auth-file", "", "Json file with auth data")
	flag.Parse()

	if *proxyAddr == "" {
		return nil, errors.New("empty proxy address")
	}
	if *socksAddr == "" {
		return nil, errors.New("empty socks address")
	}
	if !availableLogFormats[LogFormat(*logFormat)] {
		formats := slices.Collect(maps.Keys(availableLogFormats))
		return nil, fmt.Errorf("unknown log format: '%s', should be one of: %v", *logFormat, formats)
	}
	config := Config{
		ProxyAddr: *proxyAddr,
		SocksAddr: *socksAddr,
		LogLevel:  *logLevel,
		LogFormat: LogFormat(*logFormat),
		JsonAuth:  *jsonAuth,
	}
	return &config, nil
}
