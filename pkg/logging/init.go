package logging

import (
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	if log == nil {
		Init()
	}

	return log
}

func Init() {
	if log != nil {
		return
	}

	log = logrus.New()
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

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	log.Info("Successfully initialized the logrus logger")
}
