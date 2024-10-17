package connections

import (
	"ginTemplate/internal/config"
	"ginTemplate/pkg/logging"
	gormzerolog "github.com/vitaliy-art/gorm-zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	logging.DefaultLogger.Info().Msg("Connecting to database...")
	var err error
	dbDsn := config.DefaultConfig.DatabaseDsn
	logger := gormzerolog.NewGormLogger().WithInfo(func() gormzerolog.Event {
		return &gormzerolog.GormLoggerEvent{Event: logging.DefaultLogger.Info()}
	})
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger,
	})
	if err != nil {
		logging.DefaultLogger.Fatal().Msgf("Error connecting to database. Error: ", err)
	}

	return db
}
