package logger

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/exp/slog"
)

var mu sync.Mutex

// global variable for the configured logger
var logger *slog.Logger

// SetGlobalLogger sets a globally accessible logger instance
func SetGlobalLogger(lvl string) {
	mu.Lock()
	defer mu.Unlock()

	var logLevel slog.Level
	switch lvl {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}

func GetLogger() *slog.Logger {
	mu.Lock()
	defer mu.Unlock()
	return logger
}

func StructuredLog(level string, msg string, attrs ...any) {
	l := GetLogger()
	switch level {
	case "debug":
		l.Debug(msg, attrs...)
	case "info":
		l.Info(msg, attrs...)
	case "warn":
		l.Warn(msg, attrs...)
	case "error":
		l.Error(msg, attrs...)
	default:
		fmt.Println("Invalid log level:", level)
	}
}
