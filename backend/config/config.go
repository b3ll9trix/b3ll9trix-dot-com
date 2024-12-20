package config

import (
	"backend/logger"
	"io"
)

type Config struct {
	Port     string
	Domain   string
	LogLevel logger.Level
	LogFile  *io.Writer
}
