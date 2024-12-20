package cd

import (
	"backend/logger"
	"net/http"

	"backend/config"
)

func Handle(logger logger.Logger, config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
