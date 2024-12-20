package main

import (
	"backend/internal/cd"
	"backend/internal/ls"
	"backend/internal/welcome"
	"backend/logger"
	"io"
	"net/http"
)

type Config struct {
	Port     string
	Domain   string
	LogLevel logger.Level
	LogFile  *io.Writer
}

func NewServer(logger logger.Logger, config *Config) http.Handler {

	var handler http.Handler
	mux := http.NewServeMux()
	handler = mux
	addRoutes(mux, logger, config)
	// Any middlewares if any

	return handler
}

func addRoutes(mux *http.ServeMux, logger logger.Logger, config *Config) {
	mux.Handle("/api/v1/welcome", welcome.Handle(logger))
	mux.Handle("/api/v1/list", ls.Handle(logger))
	mux.Handle("/api/v1/cd", cd.Handle(logger))
}
