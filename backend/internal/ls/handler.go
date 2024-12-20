package ls

import (
	"backend/logger"
	"net/http"
)

func Handle(logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
