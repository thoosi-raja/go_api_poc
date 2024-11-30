package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitaiteLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.TraceLevel)
	log.SetOutput(os.Stdout)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Log(args ...interface{}) {
	log.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Trace(args ...interface{}) {
	log.Trace(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}
