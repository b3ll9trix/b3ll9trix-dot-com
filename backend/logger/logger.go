package logger

import (
	"io"
	"strconv"

	"github.com/rs/zerolog"
)

type Level int8

func ToLevel(str string) Level {
	levelInt, err := strconv.Atoi(str)
	if err != nil {
		return NoLevel
	}
	switch levelInt {
	case -1:
		return TraceLevel
	case 1:
		return DebugLevel
	case 2:
		return InfoLevel
	case 3:
		return WarnLevel
	case 4:
		return ErrorLevel
	case 5:
		return FatalLevel
	case 6:
		return PanicLevel

	default:
		return NoLevel
	}
}

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
)

type Logger interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Panic() *zerolog.Event
}

func New(w io.Writer) Logger {
	logger := zerolog.New(w)
	return &logger
}
