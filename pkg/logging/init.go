package logging

import (
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Init() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)

	// logging format configuration
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return f.Function, filename
		},
	}

	// print to stdout
	log.SetOutput(os.Stdout)

	logFile, err := os.OpenFile("logs/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Cannot open the logs file: %s", err.Error())
	}

	log.SetOutput(logFile)

	log.Info("Successfully initialized the logrus logger")

	return log
}
