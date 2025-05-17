package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func New(levelStr string) *Logger {
	var log *slog.Logger
	var level slog.Level

	switch levelStr {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
	)

	return &Logger{
		log: log,
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args...)
}

func (l *Logger) With(args ...interface{}) *Logger {
	return &Logger{
		log: l.log.With(args...),
	}
}
