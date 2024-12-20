package main

import (
	"backend/config"
	"backend/internal/cd"
	"backend/internal/ls"
	"backend/internal/welcome"
	"backend/logger"
	"net/http"
)

func NewServer(logger logger.Logger, config *config.Config) http.Handler {

	var handler http.Handler
	mux := http.NewServeMux()
	handler = mux
	addRoutes(mux, logger, config)
	// Any middlewares if any

	return handler
}

func addRoutes(mux *http.ServeMux, logger logger.Logger, config *config.Config) {
	mux.Handle("/api/v1/welcome", welcome.Handle(logger, config))
	mux.Handle("/api/v1/list", ls.Handle(logger, config))
	mux.Handle("/api/v1/cd", cd.Handle(logger, config))
}
