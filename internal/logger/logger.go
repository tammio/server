package logger

import (
	"log"
)

var (
	levelOrder = map[string]int{
		"debug": 0,
		"info":  1,
		"warn":  2,
		"error": 3,
	}
	currentLevel = 1
)

func SetLevel(level string) {
	if l, ok := levelOrder[level]; ok {
		currentLevel = l
	}
}

func logf(level string, format string, v ...interface{}) {
	if levelOrder[level] < currentLevel {
		return
	}
	log.Printf("["+level+"] "+format, v...)
}

func Debugf(format string, v ...interface{}) {
	logf("debug", format, v...)
}

func Infof(format string, v ...interface{}) {
	logf("info", format, v...)
}

func Warnf(format string, v ...interface{}) {
	logf("warn", format, v...)
}

func Errorf(format string, v ...interface{}) {
	logf("error", format, v...)
}
