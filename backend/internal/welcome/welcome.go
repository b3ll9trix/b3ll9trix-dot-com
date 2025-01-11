package welcome

import (
	"backend/logger"

	"github.com/a-h/templ"
)

func generateWelcome(logger logger.Logger) (templ.Component, error) {
	// Returns a html that renders the welcome html
	logger.Info().Msg("generating html for welcome.")
	return getStarthtml(), nil
}
