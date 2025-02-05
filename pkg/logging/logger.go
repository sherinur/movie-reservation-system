package logging

import (
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(env string) *Logger {
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

	if env == "test" {
		return &Logger{log}
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatalf("Cannot open the logs file: %s", err.Error())
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return &Logger{
		log,
	}
}
