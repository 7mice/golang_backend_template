package main

import (
	"ginTemplate/app"
	"ginTemplate/internal/config"
	"ginTemplate/pkg/logging"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.DefaultConfig.IsDebug {
		logging.DefaultLogger.Info().Msg("Starting application as debug")
	} else {
		gin.SetMode(gin.ReleaseMode)
		logging.DefaultLogger.Info().Msg("Starting application")
	}
	app.StartGIN()
}
