package bootstrap

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLog(config *Config) {

	switch config.LogFormat {
	case LogFormatText:
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	case LogFormatJson:
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.SetOutput(os.Stderr)

	switch config.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	}

	log.Info("Init logs")
}
