package welcome

import (
	"backend/config"
	"backend/internal/beep"
	"backend/logger"
	"net/http"
)

func Handle(logger logger.Logger, config *config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html, err := generateWelcome(logger)
		if err != nil {
			logger.Error().AnErr("err", err).Msg("error generating html for welcome.")
			logger.Info().Msg("sending server-error html.")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(beep.ServerError()))
		} else {
			logger.Info().Msg("html for welcome successfully generated.")
			logger.Debug().Str("html", html).Send()
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(html))
		}
	})
}
