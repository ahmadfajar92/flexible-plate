package log

import (
	"log/syslog"

	"github.com/sirupsen/logrus"
	logrusSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

const (
	// TOPIC for setting topic of log
	TOPIC = "scaffold-service-log"
	// LogTag default log tag
	LogTag = "scaffold-service"

	// InfoLevel log
	InfoLevel = logrus.InfoLevel

	// DebugLevel log
	DebugLevel = logrus.DebugLevel

	// ErrorLevel log
	ErrorLevel = logrus.ErrorLevel

	// FatalLevel log
	FatalLevel = logrus.FatalLevel

	// PanicLevel log
	PanicLevel = logrus.PanicLevel
)

// Context function for logging the context of echo
// c string context
// s string scope
func Context(c string, s string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

// Log function for returning entry type
// level logrus.Level
// message string message of log
// context string context of log
// scope string scope of log
func Log(level logrus.Level, message string, context string, scope string) {
	defer recover()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	syslogOutput, err := logrusSyslog.NewSyslogHook("", "", syslog.LOG_INFO, LogTag)
	logrus.AddHook(syslogOutput)

	if err != nil {
		return
	}

	entry := Context(context, scope)
	switch level {
	case logrus.DebugLevel:
		entry.Debug(message)
	case logrus.InfoLevel:
		entry.Info(message)
	case logrus.WarnLevel:
		entry.Warn(message)
	case logrus.ErrorLevel:
		entry.Error(message)
	case logrus.FatalLevel:
		entry.Fatal(message)
	case logrus.PanicLevel:
		entry.Panic(message)
	}
}
