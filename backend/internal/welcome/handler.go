package welcome

import (
	"backend/logger"
	"net/http"
)

func Handle(logger logger.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
