package config

import "github.com/sirupsen/logrus"

var log = logrus.New()

func Log() *logrus.Logger {
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return log
}
