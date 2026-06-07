package log

import (
	"io"
	"os"
	"wstore/internal/config"

	"github.com/sirupsen/logrus"
)

var file *os.File

func Init() func() error {
	config := config.GetConfig()
	var err error

	if config.DebugMode {
		logrus.SetLevel(logrus.PanicLevel)
		logrus.SetReportCaller(true)
	} else {
		level, err := logrus.ParseLevel(config.LogConfig.Level)
		if err != nil {
			logrus.Warnf("Failed to parse log level: %v", err)
			level = logrus.InfoLevel
		}
		logrus.SetLevel(level)
	}
	if config.LogConfig.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	file, err = os.OpenFile(config.LogConfig.Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logrus.WithField("path", config.LogConfig.Path).WithError(err).Errorf("Failed to open log file")
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
	}

	return func() error {
		if file != nil {
			return file.Close()
		}
		return nil
	}
}
