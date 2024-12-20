package ls

import (
	"backend/config"
	"backend/logger"
	"net/http"
)

func Handle(logger logger.Logger, config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
