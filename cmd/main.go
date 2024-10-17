package main

import (
	"ginTemplate/app"
	"ginTemplate/internal/config"
	"ginTemplate/pkg/logging"
)

func main() {
	if config.DefaultConfig.IsDebug {
		logging.DefaultLogger.Info().Msg("Starting application as debug")
	} else {
		logging.DefaultLogger.Info().Msg("Starting application")
	}
	app.StartGIN()
}
