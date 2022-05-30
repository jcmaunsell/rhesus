package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	LogDirectory   = "log"
	ServiceLogPath = LogDirectory + "/service.log"
)

var serviceLogger *logrus.Logger

func Service() *logrus.Logger {
	if err := os.MkdirAll(LogDirectory, 0700); err != nil {
		panic(fmt.Errorf("cannot create log directory: %w", err))
	}
	if serviceLogger == nil {
		serviceLogger = New(ServiceLogPath)
	}
	return serviceLogger
}

func New(outputPath string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(new(logrus.JSONFormatter))
	logFile, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Errorf("could not create log file: %w", err))
	}
	log.SetOutput(logFile)
	return log
}
