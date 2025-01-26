package utils

import (
	"log"
	"log/slog"
	"net/http"
)

type iLogger interface {
	PrintInfoMsg(mes string, args ...interface{})
	PrintDebugMsg(mes string, args ...interface{})
	PrintErrorMsg(mes string, args ...interface{})
	PrintWarnMsg(mes string, args ...interface{})
	LogRequestMiddleware(next http.Handler) http.Handler
}

type Logger struct {
	debugMode    bool
	bracketsMode bool
}

func NewLogger(debugMode bool, bracketsMode bool) iLogger {
	return &Logger{
		debugMode:    debugMode,
		bracketsMode: bracketsMode,
	}
}

func printfMsg(level string, mes string, args ...interface{}) {
	log.Printf(level+" "+mes, args...)
}

func (l *Logger) PrintInfoMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[INFO]", mes, args...)
		return
	}
	slog.Info(mes, args...)
}

func (l *Logger) PrintDebugMsg(mes string, args ...interface{}) {
	if l.debugMode {
		printfMsg("[DEBUG]", mes, args...)
	}
}

func (l *Logger) PrintErrorMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[ERROR]", mes, args...)
		return
	}
	slog.Error(mes, args...)
}

func (l *Logger) PrintWarnMsg(mes string, args ...interface{}) {
	if l.bracketsMode {
		printfMsg("[WARN]", mes, args...)
		return
	}
	slog.Warn(mes, args...)
}

func (l *Logger) LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] Request %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
